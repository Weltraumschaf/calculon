#!/usr/bin/env bash

#
# Generate code coverage for Go code.
#
# Based on https://gitlab.com/pantomath-io/demo-tools/-/blob/master/tools/coverage.sh
#

set -u

USAGE="$(basename "${0}") <coverage_dir> [html]"
COVERAGE_DIR="${1:-}"
REPORT="${2:-}"
PKG_LIST=$(go list ./...)

if [[ "${COVERAGE_DIR}" == "" ]]; then
    echo "No coverage directory given!"
    echo "${USAGE}"
    exit 1
fi

mkdir -p "${COVERAGE_DIR}"

for pkg in ${PKG_LIST}; do
    go test -covermode=count -coverprofile \
        "${COVERAGE_DIR}/${pkg##*/}.cov" "${pkg}"
done

echo 'mode: count' > "${COVERAGE_DIR}/coverage.cov"
tail -q -n +2 "${COVERAGE_DIR}"/*.cov >> "${COVERAGE_DIR}/coverage.cov"

# Display the global code coverage
go tool cover -func="${COVERAGE_DIR}/coverage.cov"

if [[ "${REPORT}" == "html" ]]; then
    go tool cover -html="${COVERAGE_DIR}/coverage.cov" -o "${COVERAGE_DIR}/coverage.html"
fi
