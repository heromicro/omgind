#!/usr/bin/env python3
"""Build a reproducible GeoNames administrative seed without downloading at startup."""
import argparse
import collections
import csv
import datetime
import gzip
import hashlib
import io
import json
from pathlib import Path
import zipfile

BASE = "https://download.geonames.org/export/dump/"


def build(raw, output):
    countries = {}
    for row in csv.reader(io.StringIO((raw / "countryInfo.txt").read_text()), delimiter="\t"):
        if not row or row[0].startswith("#") or row[0] in {"CS", "AN"}:
            continue
        countries[row[0]] = row
    nodes = {}
    keys = {}
    exclusions = collections.Counter()
    def add(code, name, ascii_name, geoid, country, level, parent):
        identifier = "gn:" + geoid
        if identifier in nodes:
            exclusions["duplicate_geoname_id"] += 1
            keys[code] = identifier
            return
        nodes[identifier] = dict(id=identifier, parent_id=parent, name=name,
            name_ascii=ascii_name, country_code=country, source_code=code,
            source="geonames", source_id=geoid, admin_level=level,
            feature_code="COUNTRY" if level == 0 else "ADM" + str(level))
        keys[code] = identifier
    for code, row in sorted(countries.items()):
        add(code, row[4], row[4], row[16], code, 0, None)
        nodes[keys[code]]["iso_alpha2"] = code if code != "XK" else None
        nodes[keys[code]]["iso_alpha3"] = row[1] if code != "XK" else None
        nodes[keys[code]]["iso_numeric"] = row[2] if code != "XK" else None
    depths = {}
    for code, row in countries.items():
        # Small countries keep primary subdivisions; no artificial empty levels.
        depths[code] = 1 if float(row[6] or 0) < 10000 and int(row[7] or 0) < 1000000 else 2
    depths.update({"CN": 3, "DE": 3})
    for level, filename in [(1, "admin1CodesASCII.txt"), (2, "admin2Codes.txt")]:
        for row in csv.reader(io.StringIO((raw / filename).read_text()), delimiter="\t"):
            if len(row) != 4:
                raise ValueError("Invalid administrative row: " + repr(row))
            code, name, ascii_name, geoid = row
            country = code.split(".")[0]
            if country not in countries or level > depths[country]:
                exclusions["outside_depth_policy"] += 1
                continue
            parent_code = code.rsplit(".", 1)[0]
            parent = keys.get(parent_code)
            if parent is None:
                exclusions["missing_parent"] += 1
                continue
            add(code, name, ascii_name, geoid, country, level, parent)
    for country in ["CN", "DE"]:
        path = raw / (country + ".zip")
        if not path.exists():
            raise ValueError("Required county-level source is missing: " + str(path))
        with zipfile.ZipFile(path) as archive:
            with archive.open(country + ".txt") as data:
                for row in csv.reader(io.TextIOWrapper(data, encoding="utf-8"), delimiter="\t"):
                    if row[6:8] != ["A", "ADM3"]:
                        continue
                    code = ".".join([country, row[10], row[11], row[12]])
                    parent = keys.get(".".join([country, row[10], row[11]]))
                    if parent is None:
                        parent = keys.get(".".join([country, row[10]]))
                        if parent is None:
                            exclusions[country + "_ADM3_missing_parent"] += 1
                            continue
                        exclusions[country + "_ADM3_missing_intermediate_level"] += 1
                    add(code, row[1], row[2], row[0], country, 3, parent)
                    nodes["gn:" + row[0]]["latitude"] = float(row[4])
                    nodes["gn:" + row[0]]["longitude"] = float(row[5])
    children = collections.defaultdict(list)
    for node in nodes.values():
        children[node["parent_id"]].append(node)
    ordered = []
    def walk(node, tree_id, path, counter):
        node["tree_id"] = tree_id
        node["tree_level"] = len(path) + 1
        node["tree_path"] = "/".join(path)
        node["tree_left"] = counter[0]; counter[0] += 1
        ordered.append(node)
        kids = sorted(children[node["id"]], key=lambda n: n["source_code"])
        node["is_leaf"] = not kids
        for child in kids:
            walk(child, tree_id, path + [node["id"]], counter)
        node["tree_right"] = counter[0]; counter[0] += 1
    for index, node in enumerate(sorted(children[None], key=lambda n: n["country_code"]), 1):
        walk(node, index, [], [1])
    output.mkdir(parents=True, exist_ok=True)
    data = (json.dumps(ordered, ensure_ascii=False, separators=(",", ":")) + "\n").encode()
    seed = gzip.compress(data, mtime=0)
    (output / "world.json.gz").write_bytes(seed)
    sources = []
    for filename in ["countryInfo.txt", "admin1CodesASCII.txt", "admin2Codes.txt", "CN.zip", "DE.zip"]:
        content = (raw / filename).read_bytes()
        sources.append(dict(url=BASE + filename, bytes=len(content), sha256=hashlib.sha256(content).hexdigest()))
    counts = collections.Counter(n["country_code"] for n in ordered)
    manifest = dict(format_version=1, collected_at=datetime.datetime.now(datetime.timezone.utc).isoformat(),
        license="CC BY 4.0", attribution="GeoNames (https://www.geonames.org/)",
        license_url="https://creativecommons.org/licenses/by/4.0/", sources=sources,
        records=len(ordered), countries=len(countries), counts_by_country=dict(sorted(counts.items())),
        max_admin_level_by_country=dict(sorted(depths.items())), exclusions=dict(exclusions),
        sha256=hashlib.sha256(seed).hexdigest(),
        notes=["GeoNames subdivision codes are not ISO 3166-2 codes.",
               "XK is a user-assigned code, not an officially assigned ISO 3166-1 code.",
               "CS and AN are historical entries and are excluded.",
               "County equivalence varies by country; coverage reflects source records, not a completeness guarantee.",
               "Names are source names; ASCII names are transliterations, not guaranteed English translations."])
    (output / "manifest.json").write_text(json.dumps(manifest, ensure_ascii=False, indent=2) + "\n")
    print(json.dumps(dict(records=len(ordered), countries=len(countries), exclusions=dict(exclusions))))

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--raw-dir", type=Path, required=True)
    parser.add_argument("--output", type=Path, default=Path("configs/data/districts"))
    args = parser.parse_args()
    build(args.raw_dir, args.output)
