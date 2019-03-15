SOURCES=./

dep:
	dep ensure

non_panicking_example:
	go run ${SOURCES}/examples/non_panicking/main.go

panicking_example:
	go run ${SOURCES}/examples/panicking/main.go

examples: panicking_example non_panicking_example

.DEFAULT_GOAL := test
.PHONY: test
test: examples

