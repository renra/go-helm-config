SOURCES=./
BINS=bin

dep:
	dep ensure

non_panicking_example:
	go run ${SOURCES}/examples/non_panicking/main.go

panicking_example:
	go run ${SOURCES}/examples/panicking/main.go

examples: panicking_example non_panicking_example

.PHONY: clean_tests
clean_tests:
	rm -rf ${BINS}/tests

.PHONY: build_tests
build_tests: clean_tests
	go test ./test... -count 1 -v -c -o ${BINS}/tests

.PHONY: test
.DEFAULT_GOAL := test
test: build_tests
	${BINS}/tests
