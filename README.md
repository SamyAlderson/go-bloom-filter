# go-bloom-filter

A Go implementation of a Bloom filter with counting.

## What and Why

A Bloom filter is a space-efficient probabilistic data structure that is used to test whether an element is a member of a set. This implementation includes a counting Bloom filter, which keeps track of the number of occurrences of each element.

## Install

Run `go get` to install the package:
```bash
go get github.com/samyalderson/go-bloom-filter
```
## Usage

```go
import (
	"fmt"

	"github.com/samyalderson/go-bloom-filter"
)

func main() {
	// Create a new Bloom filter with a capacity of 100 and a hash count of 5
	bf := bloom.New(100, 5)

	// Add some elements to the filter
	bf.Add("hello")
	bf.Add("world")
	bf.Add("go")

	// Test whether an element is in the filter
	if bf.Test("hello") {
		fmt.Println("hello is in the filter")
	}

	// Create a new counting Bloom filter with a capacity of 100 and a hash count of 5
	cbf := count.New(100, 5)

	// Add some elements to the filter
	cbf.Add("hello")
	cbf.Add("world")
	cbf.Add("go")

	// Get the count of an element
	count, ok := cbf.Get("hello")
	if ok {
		fmt.Printf("hello has been added %d times\n", count)
	}
}
```
## Build from Source

Run `go build` to build the package:
```bash
go build main.go
```
## Project Structure

* `go.mod`: Go module file
* `go.sum`: Go checksum file
* `main.go`: Main entry point
* `bloom.go`: Bloom filter implementation
* `count.go`: Counting Bloom filter implementation
* `test_bloom_test.go`: Bloom filter test suite
* `test_count_test.go`: Counting Bloom filter test suite
* `Makefile`: Build script
* `README.md`: Project README
* `.gitignore`: Git ignore file

## License

This project is licensed under the MIT License.

## Credits

This project was inspired by the `github.com/olahol/bloom-filter` package.

## Tests

The package includes test coverage for the Bloom filter and counting Bloom filter implementations. Run `go test` to run the tests:
```bash
go test
```
## Dependencies

* `github.com/stretchr/testify/assert`: Test package for Go.