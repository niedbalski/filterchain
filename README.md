# Go-filterchain

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/niedbalski/filterchain)
[![Build Status](https://travis-ci.org/niedbalski/filterchain.svg?branch=master)](https://travis-ci.org/niedbalski/filterchain)

very simple implementation of a filter chain pattern written in Golang.

### Usage

```go
package main

import (
	"github.com/niedbalski/filterchain"
	"strings"
	"fmt"
)

type TestToLowerFilter struct {}
type TestRepeatFilter struct {}

func (filter TestToLowerFilter) Apply(t interface{}) (interface{}) {
	return strings.ToLower(t.(string))
}

func (filter TestRepeatFilter) Apply(t interface{}) (interface{}) {
	return strings.Repeat(t.(string), 3)
}

func main() {
	fmt.Println(filterchain.New(TestToLowerFilter{}, TestRepeatFilter{}).Apply("test"))
}

```
String will be produced as:

```shell
testtestest
```

