#!/usr/bin/env bash

set -o nounset
set -o noclobber
set -o errexit
set -o pipefail

echo "initializing ..."

#[[ -f go.mod ]] || \
#docker run -ti --rm \
#-e HOME=${HOME} \
#-v "${HOME}:${HOME}/" \
#-v /etc/group:/etc/group:ro \
#-v /etc/passwd:/etc/passwd:ro \
#-u $(id -u ${USER}):$(id -g ${USER}) \
#-w $(pwd) \
#golang:1.14.4-stretch go mod init "infra_test"

echo "formatting ..."

docker run -ti --rm \
-e HOME=${HOME} \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
golang:1.14.4-stretch go fmt

echo "testing ..."

mkdir -p $(pwd)/.terraform/test-report

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e TERRATEST_IAM_ROLE=arn:aws:iam::358458405859:role/terratest \
-e GOMAXPROCS=5 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-v $(which terraform):/usr/local/bin/terraform:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
golang:1.14.4-stretch go test -v -timeout 30m -count=1 | tee $(pwd)/.terraform/test-report.log

terratest_log_parser -testlog $(pwd)/.terraform/test-report.log -outputdir $(pwd)/.terraform/test-report
