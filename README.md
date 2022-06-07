# Arctic Tern
[![Build Status](https://app.travis-ci.com/arikama/go-arctic-tern.svg?branch=master)](https://app.travis-ci.com/arikama/go-arctic-tern)
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
	"github.com/arikama/go-mysql-test-container/mysqltestcontainer"
)

func main() {
	db, _ := mysqltestcontainer.Start("test", "")
	migrationDir := "./migration/example"
	arctictern.Migrate(db, migrationDir)
}

```
