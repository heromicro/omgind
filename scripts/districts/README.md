# Global administrative seed

The source is GeoNames, licensed under CC BY 4.0. Attribution: GeoNames, https://www.geonames.org/. The original source readme is retained in `configs/data/districts/GEONAMES-README.txt`.

## Rebuild

Download the following files from `https://download.geonames.org/export/dump/` into a scratch directory:

- `countryInfo.txt`
- `admin1CodesASCII.txt`
- `admin2Codes.txt`
- `CN.zip`
- `DE.zip`

Use `curl -fL -C -` for resumable archive downloads. Then run from the project root:

```sh
python3 scripts/districts/collect.py --raw-dir /tmp/omgind-geonames
```

The script emits a compressed JSON snapshot and a manifest containing source URLs, SHA-256 hashes, record counts, country depth policy and exclusions. Raw archives are not required by the running application and are not committed.

## Scope and interpretation

Country/territory records retain source-provided ISO country codes; XK is explicitly marked as a user-assigned code. Historical CS and AN entries are excluded. Administrative subdivision identifiers are GeoNames identifiers, not ISO 3166-2 codes and not national statistical codes. The administrative tree is a source-derived hierarchy, not a statement about sovereignty or boundaries.

Most countries retain ADM1 and ADM2. Countries/territories smaller than 10,000 square kilometers with fewer than one million inhabitants retain ADM1 only; this explicit size heuristic keeps small-country trees useful without inventing empty levels. China and Germany additionally retain ADM3 to reach county/district equivalents. No ordinary populated places, historical divisions, villages or ADM4/ADM5 records are included. County equivalence varies by country; ADM2 is not universally a county, and the snapshot does not guarantee exhaustive or legally authoritative coverage.

For ADM3 records with no matching ADM2 entry, the nearest available ADM1 is used and the manifest records that the intermediate level is missing. Records with no usable parent are excluded and counted. Source names are retained. ASCII transliterations are preserved in the file but are not mislabeled as English translations. This release does not supply a verified Chinese localization for every global place name.

## Startup import

`[District]` in each application config controls startup initialization. The startup reads only `configs/data/districts/world.json.gz`; it never contacts GeoNames. Rows use stable `gn:<geonameid>` IDs, preserve parent links, and include precomputed nested-set tree bounds. The database `extra` field stores the GeoNames source ID; `st_code` is not populated with incompatible GeoNames codes.

Initialization imports only into an entirely empty `sys_districts` table, in a transaction with batches. An existing database, including the older China-only dataset, is left untouched to avoid duplicate countries, mixed hierarchies and broken address references. Importing the global snapshot over an existing installation requires an explicit migration of existing district references. Do not delete production district rows just to trigger initialization.

Use a single database migration/startup job for the initial import. Multiple simultaneous initializers may race; a conflicting transaction fails and rolls back rather than importing a partial tree. Re-running after a successful initialization skips the import. Updates to an existing snapshot require a separately reviewed migration.

No live database is modified during data collection or tests. SQLite tests validate the file and import behavior.
