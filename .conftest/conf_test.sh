#!/usr/bin/env bash

set -o nounset
set -o noclobber
set -o errexit
set -o pipefail

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
instrumenta/conftest --version

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
instrumenta/conftest parse ../*.tf --combine --no-color >| module.json

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
instrumenta/conftest test --no-color --combine --policy module.rego ../*.tf

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e TF_INPUT=0 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd)/../.infratest \
hashicorp/terraform:0.12.28 init

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e TF_INPUT=0 \
-e AWS_REGION=us-east-1 \
-e TF_VAR_region=us-east-1 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd)/../.infratest \
hashicorp/terraform:0.12.28 plan -no-color -out=module.tfplan -parallelism=25

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-e TF_INPUT=0 \
-e AWS_REGION=us-east-1 \
-e TF_VAR_region=us-east-1 \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd)/../.infratest \
hashicorp/terraform:0.12.28 show -no-color -json module.tfplan >| tfplan.json

docker run -ti --rm \
-e HOME=${HOME} \
-e AWS_PROFILE=terraform-infra-test \
-v "${HOME}:${HOME}/" \
-v /etc/group:/etc/group:ro \
-v /etc/passwd:/etc/passwd:ro \
-u $(id -u ${USER}):$(id -g ${USER}) \
-w $(pwd) \
instrumenta/conftest test --input json --policy plan.rego tfplan.json

