FIXTURES=$(shell find * -maxdepth 0 -type d)

.PHONY: $(FIXTURES)
$(FIXTURES):
	cd $@/ && terraform init
	cd $@/ && terraform plan -out=plan.tfplan
	cd $@/ && terraform show -json plan.tfplan > plan.json

.PHONY: generate
generate: $(FIXTURES)
