# Determine this makefile's path.
# Be sure to place this BEFORE `include` directives, if any.
# THIS_FILE := $(lastword $(MAKEFILE_LIST))
THIS_FILE := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

include vars.mk

#################################
# Bazel targets
#################################

.PHONY: yarn
yarn: ## install dependencies
	@bazel run @nodejs//:yarn

.PHONY: build
build: yarn ## build all targets
	@bazel run @nodejs//:yarn -- run build
	@bazel build //...

.PHONY: test
test: yarn ## test all targets
	@bazel run @nodejs//:yarn -- run build
	@bazel test //...

.PHONY: buildifier
buildifier: ## format Bazel WORKSPACE and BUILD.bazel files
	@bazel run :buildifier

.PHONY: gen-go-deps
gen-go-deps: ## regenerate dependencies file (deps.bzl)
	@bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies

.PHONY: docker-publish
docker-publish: metadata-docker-publish ## publish all docker images

#################################
# Admin service targets
#################################

.PHONY: admin-docker-build
admin-docker-build: build ## publish linux/amd64 platform image locally
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/admin/backend/cmd:image -- --norun

.PHONY: admin-docker-publish
admin-docker-publish: build ## publish linux/amd64 platform image to Dockerhub
	@bazel run --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 //services/admin/docker:push

.PHONY: admin-service-run
admin-service-run: build ## start the admin-service
	@bazel run //services/admin/backend/cmd

.PHONY: admin-service-test
admin-service-test: ## run all admin-service tests
	@bazel test //services/admin/backend/...

#################################
# API-SPA-Server targets
#################################

.PHONY: apispa-test
apispa-test: ## run API-SPA-Server tests
	@bazel test //shared/go/pkg/server/...

#################################
# Docs targets
#################################

.PHONY: docs-build
docs-build: build ## build the DSP API Slate docs
	docker run --rm --name slate -v $(CURRENT_DIR)/docs:/srv/slate/source slatedocs/slate build

.PHONY: docs-serve
docs-serve: ## serve the DSP API Slate docs locally
	docker run --rm --name slate -p 4567:4567 -v $(CURRENT_DIR)/docs:/srv/slate/source slatedocs/slate serve

.PHONY: docs-publish
docs-publish: publish ## publish the DSP API Slate docs to Github Pages
	docker run --rm --name slate -v $(CURRENT_DIR)/docs:/srv/slate/source slatedocs/slate publish

#################################
# Other targets
#################################

.PHONY: help
help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
