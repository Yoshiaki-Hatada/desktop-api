TEST_METRICS_SOCKET=/tmp/docker-cli.sock

GREEN=\033[0;32m
NC=\033[0m # No Color

.PHONY: all
all: generate

.PHONY: gen-desktop-cli
gen-desktop-cli:
	rm -rf ./internal/generated/docker-cli
	openapi-generator-cli generate -g go -i ./docker-cli-api.yaml \
		-o ./internal/generated/docker-cli \
		--additional-properties=packageName=dockercliapi \
		--additional-properties=enumClassPrefix=true \
		--additional-properties=generateInterfaces=true \
		--additional-properties=isGoSubmodule=true
	openapi-generator-cli generate -g html2 -i ./docker-cli-api.yaml \
		-o ./internal/generated/docker-cli/html
	goimports -w .
	rm -rf ./internal/generated/docker-cli/.gitignore \
		./internal/generated/docker-cli/.openapi-generator-ignore \
		./internal/generated/docker-cli/.travis.yml \
		./internal/generated/docker-cli/git_push.sh \
		./internal/generated/docker-cli/go.mod \
		./internal/generated/docker-cli/go.sum

.PHONY: generate
generate: gen-desktop-cli

.PHONY: test
test: test-unit test-e2e
	@echo -e "All tests ${GREEN}PASSED${NC}!"

.PHONY: test-unit
test-unit:
	go test -cover $(shell go list ./... | grep -vE 'e2e')

.PHONY: test-e2e
test-e2e:
	rm -f $(TEST_METRICS_SOCKET)
	TEST_METRICS_SOCKET=$(TEST_METRICS_SOCKET) go test -tags e2e -cover $(shell go list -tags e2e ./... | grep 'e2e')
