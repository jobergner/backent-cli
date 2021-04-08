package main

import (
	"bar-cli/enginefactory"
	"bar-cli/examples"
	"fmt"
)

func main() {
	state := enginefactory.WriteEngineFrom(examples.StateConfig)
	fmt.Println(string(state))
}
