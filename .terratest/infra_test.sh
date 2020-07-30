#!/usr/bin/env bash

set -o nounset
set -o noclobber
set -o errexit
set -o pipefail


echo "initializing ..."

export SCRIPTS_GIST=https://gist.githubusercontent.com/btower-labz/f9f60b8321307154c56dc75cb4c293f6/raw

curl --location --silent ${SCRIPTS_GIST}/golang-docker-init-module.sh | /bin/sh

curl --location --silent ${SCRIPTS_GIST}/golang-docker-format-module.sh | /bin/sh

curl --location --silent ${SCRIPTS_GIST}/golang-docker-test-module.sh | /bin/sh | tee $(pwd)/report.log

test -f terratest_log_parser && terratest_log_parser -testlog $(pwd)/report.log -outputdir $(pwd)/report
