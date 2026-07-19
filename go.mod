go.mod
```go
module go-bloom-filter

go 1.18

require (
	github.com/stretchr/testify/assert v1.8.0
)

replace (
	github.com/stretchr/testify/assert => file:///gopath/src/github.com/stretchr/testify/assert v1.8.0
)

require (
	github.com/philhofer/fwd v0.0.0-20160220053058-7df1b5f4a5b5
)
```