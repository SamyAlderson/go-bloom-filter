# Build script for go-bloom-filter

# Set up build flags
GOFLAGS := -trimpath -mod=readonly

# Set up test flags
TESTFLAGS := -coverprofile=coverage.out -covermode=atomic

# Set up dependencies
TESTDEPENDENCIES := github.com/stretchr/testify/assert

# Build the project
build:
	@go build $(GOFLAGS) -o go-bloom-filter main.go

# Test the project
test:
	@go test $(TESTFLAGS) ./...

# Run the project
run:
	@go run main.go

# Clean the project
clean:
	@rm -f go-bloom-filter main.go

# Get dependencies
get:
	@go mod download

# Format the code
format:
	@goimports -w .
	@gofmt -s -d .

# Generate test coverage report
coverage:
	@go tool cover -func=coverage.out

# List dependencies
deps:
	@go mod graph

# Print this message
help:
	@echo "make build  Build the project"
	@echo "make test   Test the project"
	@echo "make run    Run the project"
	@echo "make clean  Clean the project"
	@echo "make get    Get dependencies"
	@echo "make format Format the code"
	@echo "make coverage Generate test coverage report"
	@echo "make deps    List dependencies"