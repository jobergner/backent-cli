package main

import (
	"bar-cli/getstartedfactory"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const outDir = "./tmp"
const outFile = "state.go"

var configNameFlag = flag.String("config", "./barcli.config.json", "path of config")

func main() {

	c, err := readConfig()
	if err != nil {
		panic(err)
	}

	validationErrs := validateConfig(c)
	if len(validationErrs) != 0 {

		for _, validationErr := range validationErrs {
			fmt.Println(validationErr)
		}

		panic("errors during config validation")
	}

	code := writeCode(c)
	if err := ioutil.WriteFile(filepath.Join(outDir, outFile), code, 0644); err != nil {
		panic(err)
	}

	if err := generateMarshallers(); err != nil {
		panic(err)
	}

	fmt.Println(getstartedfactory.WriteGetStarted(c.State, c.Actions))
}
