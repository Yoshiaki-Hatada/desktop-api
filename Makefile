TEST_METRICS_SOCKET=/tmp/docker-cli.sock

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
	rm -rf ./internal/generated/docker-cli/.gitignore \
		./internal/generated/docker-cli/.openapi-generator-ignore \
		./internal/generated/docker-cli/.travis.yml \
		./internal/generated/docker-cli/git_push.sh \
		./internal/generated/docker-cli/go.mod \
		./internal/generated/docker-cli/go.sum

.PHONY: generate
generate: gen-desktop-cli

.PHONY: test
test:
	rm -f $(TEST_METRICS_SOCKET)
	TEST_METRICS_SOCKET=$(TEST_METRICS_SOCKET) go test -tags e2e -cover $(shell go list -tags e2e ./... | grep -vE 'e2e')
