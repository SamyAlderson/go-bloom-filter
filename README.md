# go-bloom-filter
A Go implementation of a Bloom filter with counting for approximate set membership testing.

## What it does
go-bloom-filter is a simple, efficient Bloom filter with counting that allows for approximate set membership testing. I wrote it because Go's standard library doesn't have a built-in Bloom filter implementation.

## Installation
To use go-bloom-filter, run:
```bash
go get github.com/samyalder/CS_theory/go-bloom-filter
```
## Usage
```go
import (
	"fmt"
	"github.com/samyalder/CS_theory/go-bloom-filter"
)

func main() {
	bf := go_bloom_filter.New(100, 0.01)
	bf.Add("apple")
	bf.Add("banana")
	if bf.Test("apple") {
		fmt.Println("apple is probably in the set")
	}
	if !bf.Test("cherry") {
		fmt.Println("cherry is probably not in the set")
	}
}
```
## Building from source
Run:
```bash
go build
```
## Tests
Run:
```bash
go test
```
## Project structure
* `go_bloom_filter.go`: main implementation
* `test.go`: test suite
* `main.go`: example usage
* `bf.go`: Bloom filter with counting
* `hash.go`: hash functions
* `bits.go`: bit manipulation utilities

## License
MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.