TESTS_RESULTS_FOLDER:=build/tests
TOOLS_FOLDER:=tools
GOLANGCI_LINT_EXECUTABLE:="$(TOOLS_FOLDER)/golangci-lint"
GOLANGCI_LINT_VERSION_REQUIRED="$(shell cat .golangci.version)"

.PHONY: install
install:
	go mod tidy

tools/golangci-lint:
	@echo "Installing golangci-lint v$(GOLANGCI_LINT_VERSION_REQUIRED)..."
	mkdir -p $(TOOLS_FOLDER)
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(TOOLS_FOLDER) $(GOLANGCI_LINT_VERSION_REQUIRED)
	$(GOLANGCI_LINT_EXECUTABLE) version

.PHONY: lint-fix
lint-fix: tools/golangci-lint
	$(GOLANGCI_LINT_EXECUTABLE) run --fix

.PHONY: lint
lint: tools/golangci-lint
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
