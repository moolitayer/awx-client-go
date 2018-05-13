all: verify check build examples
.PHONY: all

verify:
	hack/verify.sh
.PHONY: verify

check: verify
	go test ./awx
.PHONY: check

build:
	go build ./awx
.PHONY: build

examples: build
	hack/build-examples.sh
.PHONY: examples
