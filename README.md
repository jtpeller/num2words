# num2words

## Overview

Converts a number to its English counterpart 

## Usage

Installation can be achieved via:

```
go get -u github.com/jtpeller/num2words
```

Then, import it:

```go
import "github.com/jtpeller/num2words"
```

You can test the package via `go test ./num2words_test`

## Contents

- `test`
    - `test.go` -- holds all test functions for the package
- `bignum.go` -- holds needed wrappers for the big number package. Based on my repo at [gobig](https://github.com/jtpeller/gobig)
- `go.mod` -- manages the module
- `go.sum` -- manages any dependencies (which itself is managed by `go.mod`)
- `num2words.go` -- implementation for converting numbers to their word representation (e.g. 5 => five)
- `README.md` -- the file you're reading

