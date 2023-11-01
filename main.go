package main

import (
	"fmt"

	core "github.com/Leakageonthelamp/go-leakage-core"
)

func main() {
	test := core.NewEnv().Config()
	fmt.Println(test)
}
