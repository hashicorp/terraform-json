DOCKER_IMAGE_TAG=0.12.0-beta2
DOCKER_IMAGE="hashicorp/terraform:${DOCKER_IMAGE_TAG}"
TERRAFORM=docker run --env AWS_ACCESS_KEY_ID --env AWS_SECRET_ACCESS_KEY --env AWS_SESSION_TOKEN --rm --workdir "$$(pwd)" --volume "$$(pwd)":"$$(pwd)" $(DOCKER_IMAGE)

FIXTURES ?= $(shell find * -maxdepth 0 -type d  ! -name provider_version ! -name state_depends_on)

.PHONY: $(FIXTURES)
$(FIXTURES):
	cd $@/ && $(TERRAFORM) init
	cd $@/ && $(TERRAFORM) plan -out=plan.tfplan
	cd $@/ && $(TERRAFORM) show -json plan.tfplan > plan.json
	cd $@/ && $(TERRAFORM) providers schema -json > schemas.json

.PHONY: generate
generate: $(FIXTURES)
