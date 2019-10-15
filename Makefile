# Makefile includes some useful commands to build or format incentives
# More commands could be added

# Variables
PROJECT = gouberfmt
REPO_ROOT = github.com/gophersbd
DOCKER_PROJECT = sadlil
ROOT = ${REPO_ROOT}/${PROJECT}

fmt:
	goimports -w . analyzers
	gofmt -s -w . analyzers

compile: fmt
	go install -mod=vendor .

check: fmt
	golangci-lint run --deadline 10m ./...
	staticcheck -checks="all,-S1*" ./...

dep:
	go mod download
	go mod vendor
	go mod tidy

build:
	CGO_ENABLED=0 go build -mod=vendor -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/gouberfmt ./...

docker.check:
	docker run -i --rm \
		-v $(PWD):/modules \
		sadlil/gobuild:1.12.9 \
		make check

docker.test:
	docker run -i --rm \
		-v $(PWD):/modules $(BUILD_ARGS) \
		sadlil/gobuild:1.12.9 \
		make test

# A user can invoke tests in different ways:
#  - make test runs all tests;
#  - make test TEST_TIMEOUT=10 runs all tests with a timeout of 10 seconds;
#  - make test TEST_PKG=./model/... only runs tests for the model package;
#  - make test TEST_ARGS="-v -short" runs tests with the specified arguments;
#  - make test-race runs tests with race detector enabled.
TEST_TIMEOUT = 60
TEST_PKGS ?= ./...
TEST_TARGETS := test-short test-verbose test-race test-cover
.PHONY: $(TEST_TARGETS) test
test-short:   TEST_ARGS=-short
test-verbose: TEST_ARGS=-v
test-race:    TEST_ARGS=-race
test-cover:   TEST_ARGS=-cover
$(TEST_TARGETS): test

test: compile
	go test -mod=vendor -timeout $(TEST_TIMEOUT)s $(TEST_ARGS) $(TEST_PKGS)

clean:
	@rm -rf bin
	@rm -rf builds
	@go clean

VERSION ?= ''
release:
	@./hack/scripts/quick_release.sh $(VERSION)
	git push --tag

.PHONY: help
help:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | \
		awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | \
		sort | \
		egrep -v -e '^[^[:alnum:]]' -e '^$@$$'