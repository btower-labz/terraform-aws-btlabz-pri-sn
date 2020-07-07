#!/usr/bin/env bash

set -o nounset
set -o noclobber
set -o errexit
set -o pipefail

echo "initializing ..."

[[ -f go.mod ]] || \
docker run -ti --rm \
-e HOME=${HOME} \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
golang:1.14.4-stretch go mod init "infra_test"

echo "testing ..."

docker run -ti --rm \
-e HOME=${HOME} \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
golang:1.14.4-stretch go fmt

echo "testing ..."

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e GOMAXPROCS=5 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-v $(which terraform):/usr/local/bin/terraform:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
golang:1.14.4-stretch go test -v -timeout 30m -count=1 | tee test_out.log

terratest_log_parser -testlog test_out.log -outputdir test_out
