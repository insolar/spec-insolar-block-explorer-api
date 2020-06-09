#!/usr/bin/env bash
set -v
set -e

SCRIPT_WORKDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
PLATFORM_DIR=${SCRIPT_WORKDIR}/..

docker build --pull --rm --no-cache  -t block-explorer-spec-server -f ${SCRIPT_WORKDIR}/server.dockerfile "${PLATFORM_DIR}/client"

rm -rf "${PLATFORM_DIR}/server/" || true
mkdir -p "${PLATFORM_DIR}/server/"
id=$(docker create block-explorer-spec-server)
docker cp $id:/server/. "${PLATFORM_DIR}/server"
docker rm $id
