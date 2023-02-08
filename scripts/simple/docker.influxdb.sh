#!/usr/bin/env bash

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
echo "${BASE_DIR}"
SCRIPT_DIR=$(dirname "${BASE_DIR}")
PROJECT_DIR=$(dirname "${SCRIPT_DIR}")

docker run -d -p 8086:8086 \
      -v ${PROJECT_DIR}/data/influxdb/data:/var/lib/influxdb2 \
      -v ${PROJECT_DIR}/data/influxdb/config:/etc/influxdb2 \
      -e DOCKER_INFLUXDB_INIT_MODE=setup \
      -e DOCKER_INFLUXDB_INIT_USERNAME="omgind" \
      -e DOCKER_INFLUXDB_INIT_PASSWORD="123456" \
      -e DOCKER_INFLUXDB_INIT_ORG="omgind" \
      -e DOCKER_INFLUXDB_INIT_BUCKET="omgind" \
      -e DOCKER_INFLUXDB_INIT_RETENTION=1w \
      -e DOCKER_INFLUXDB_INIT_ADMIN_TOKEN="123456" \
      --name influxdb \
      influxdb:2.0-alpine
