TESTS_RESULTS_FOLDER:=build/tests
GOLANGCI_LINT_EXECUTABLE:="$(shell go env GOPATH)/bin/golangci-lint"

.PHONY: install
install:
	go mod tidy

.PHONY: lint-fix
lint-fix:
	$(GOLANGCI_LINT_EXECUTABLE) run --fix

.PHONY: lint
lint:
	$(GOLANGCI_LINT_EXECUTABLE) run

.PHONY: tests
tests:
	go test ./...

.PHONY: tests-coverage
tests-coverage:
	mkdir -p "$(TESTS_RESULTS_FOLDER)"
	go test -cover -coverprofile="$(TESTS_RESULTS_FOLDER)/coverage.out" ./...

$(TESTS_RESULTS_FOLDER)/coverage.out: tests-coverage

.PHONY: tests-coverage-display
tests-coverage-display: $(TESTS_RESULTS_FOLDER)/coverage.out
	go tool cover -html=build/tests/coverage.out
