DOCKER_IMAGE_NAME=terraform-0.12-local
TERRAFORM=docker run --env AWS_ACCESS_KEY_ID --env AWS_SECRET_ACCESS_KEY --env AWS_SESSION_TOKEN --rm --workdir "$$(pwd)" --volume "$$(pwd)":"$$(pwd)" $(DOCKER_IMAGE_NAME)

FIXTURES ?= $(shell find * -maxdepth 0 -type d  ! -name provider_version)

.PHONY: build-terraform-0.12-local
build-terraform-0.12-local:
	docker build --tag $(DOCKER_IMAGE_NAME) ../build-support/docker-0.12-prerelease

.PHONY: $(FIXTURES)
$(FIXTURES): build-terraform-0.12-local
	cd $@/ && $(TERRAFORM) init
	cd $@/ && $(TERRAFORM) plan -out=plan.tfplan
	cd $@/ && $(TERRAFORM) show -json plan.tfplan > plan.json

.PHONY: generate
generate: $(FIXTURES)
