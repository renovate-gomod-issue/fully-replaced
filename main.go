package main

import (
	"fmt"

	"github.com/renovate-gomod-issue/bar"
	"github.com/renovate-gomod-issue/foo"
)

func main() {
	fmt.Println(foo.Foo())
	fmt.Println(bar.Bar())
}
