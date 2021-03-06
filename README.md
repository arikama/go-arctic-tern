# Arctic Tern
[![build](https://github.com/arikama/go-arctic-tern/actions/workflows/build.yaml/badge.svg)](https://github.com/arikama/go-arctic-tern/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/arikama/go-arctic-tern/branch/master/graph/badge.svg?token=xTbyaIEFCN)](https://codecov.io/gh/arikama/go-arctic-tern)

![Arctic Tern](./arctic_tern.jpeg)

[Arctic Tern](https://en.wikipedia.org/wiki/Arctic_tern#:~:text=The%20Arctic%20tern%20is%20famous%20for%20its%20migration) is a small MySQL migration library for [Golang](https://golang.org/).

## Usage

Execute

```
go get github.com/arikama/go-arctic-tern/arctictern
```

Code

```go
package main

import (
	"github.com/arikama/go-arctic-tern/arctictern"
)

func main() {
	arctictern.Migrate(db, "./migration/example")
	arctictern.Seed(db, "./seed/example")
}
```
