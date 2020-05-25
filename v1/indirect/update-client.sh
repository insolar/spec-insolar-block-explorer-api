#!/usr/bin/env bash
set -v
set -e

SCRIPT_WORKDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
PLATFORM_DIR=${SCRIPT_WORKDIR}/..

docker build --pull --rm --no-cache  -t block-explorer-spec -f ${SCRIPT_WORKDIR}/client.dockerfile "${PLATFORM_DIR}/openapi"

rm -rf "${PLATFORM_DIR}/client/"
id=$(docker create block-explorer-spec)
docker cp $id:/spec/. "${PLATFORM_DIR}/client"
docker cp $id:/html/. "${PLATFORM_DIR}/html"
docker rm $id
