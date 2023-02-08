#!/usr/bin/env bash
BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
echo "${BASE_DIR}"
SCRIPT_DIR=$(dirname "${BASE_DIR}")
PROJECT_DIR=$(dirname "${SCRIPT_DIR}")

echo "${PROJECT_DIR}"
echo "${PWD}"

exit

docker run -d -p 15691:15691 -p 15692:15692 -p 25672:25672 -p 4369:4369 -p 5671:5671 -p 5672:5672 \
      -v ${PROJECT_DIR}/data/rabbitmq/data:/var/lib/rabbitm \
      -e RABBITMQ_DEFAULT_USER="omgind" \
      -e RABBITMQ_DEFAULT_PASS="123456" \
      --name rabbitmq
      rabbitmq:alpine
