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
	@bazel build //...

.PHONY: build-linux
build-linux: yarn ## build all targets
	@bazel build --platforms=@build_bazel_rules_nodejs//toolchains/node:linux_amd64 //...

#################################
# Docker targets
#################################

.PHONY: docker-build
docker-build: yarn ## publish linux/amd64 platform image locally
	@bazel run --platforms=@build_bazel_rules_nodejs//toolchains/node:linux_amd64 //docker -- --norun

.PHONY: docker-publish
docker-publish: yarn ## publish linux/amd64 platform image to Dockerhub
	@bazel run --platforms=@build_bazel_rules_nodejs//toolchains/node:linux_amd64 //docker:push

.PHONY: help
help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
