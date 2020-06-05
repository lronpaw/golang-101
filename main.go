package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("hello blah"))
}

func Bar() string {
	return "bar"
}

func Foo() string {
	return "foo"
}

func Qux(v string) string {
	if v == "foo" {
		return Foo()
	}

	if v == "bar" {
		return Bar()
	}

	return "INVALID"
}
