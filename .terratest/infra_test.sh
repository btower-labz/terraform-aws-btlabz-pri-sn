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
btowerlabz/docker-cloudbuild-terratest:latest go mod init "infra_test"

echo "formatting ..."

docker run -ti --rm \
-e HOME=${HOME} \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
btowerlabz/docker-cloudbuild-terratest:latest go fmt "infra_test"

echo "testing ..."

mkdir -p $(pwd)/.terraform/test-report

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e TERRATEST_IAM_ROLE=arn:aws:iam::358458405859:role/terratest \
-e GOMAXPROCS=5 \
-e GO111MODULE=on \
-e TERRATEST_REGION=us-east-1 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-v $(which terraform):/usr/local/bin/terraform:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
btowerlabz/docker-cloudbuild-terratest:latest go test -v -timeout 30m -count=1 | tee $(pwd)/.terraform/test-report.log

terratest_log_parser -testlog $(pwd)/.terraform/test-report.log -outputdir $(pwd)/.terraform/test-report
