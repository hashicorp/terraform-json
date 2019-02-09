DOCKER_IMAGE_NAME=terraform-0.12-local
TERRAFORM=docker run --rm --workdir "$$(pwd)" --volume "$$(pwd)":"$$(pwd)" $(DOCKER_IMAGE_NAME)

FIXTURES ?= $(shell find * -maxdepth 0 -type d)

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
