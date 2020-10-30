AWS_ENV_VARS=AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY AWS_SESSION_TOKEN AWS_REGION
TERRAFORM_DOCKER_CMD=docker run $(foreach var,$(AWS_ENV_VARS),--env $(var)) --rm --workdir "$$(pwd)" --volume "$$(pwd)":"$$(pwd)"

TERRAFORM_012_VERSION=0.12.11
TERRAFORM012=$(TERRAFORM_DOCKER_CMD) docker.mirror.hashicorp.services/hashicorp/terraform:$(TERRAFORM_012_VERSION)

TERRAFORM_013_VERSION=0.13.1
TERRAFORM013=$(TERRAFORM_DOCKER_CMD) docker.mirror.hashicorp.services/hashicorp/terraform:$(TERRAFORM_013_VERSION)

FIXTURES ?= $(shell find * -maxdepth 0 -type d -not -name "013*")
013FIXTURES ?= $(shell find * -maxdepth 0 -type d -name "013*")

.PHONY: $(FIXTURES)
$(FIXTURES):
	cd $@/ && $(TERRAFORM012) init
	cd $@/ && $(TERRAFORM012) plan -out=plan.tfplan
	cd $@/ && $(TERRAFORM012) show -json plan.tfplan > plan.json
	cd $@/ && $(TERRAFORM012) providers schema -json > schemas.json

.PHONY: $(013FIXTURES)
$(013FIXTURES):
	cd $@/ && $(TERRAFORM013) init
	cd $@/ && $(TERRAFORM013) plan -out=plan.tfplan
	cd $@/ && $(TERRAFORM013) show -json plan.tfplan > plan.json
	cd $@/ && $(TERRAFORM013) providers schema -json > schemas.json

.PHONY: generate
generate: $(FIXTURES) $(013FIXTURES)
