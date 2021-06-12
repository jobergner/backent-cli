package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/Java-Jonas/bar-cli/getstartedfactory"
)

const outFile = "state.go"

var configNameFlag = flag.String("config", "./barcli.config.json", "path of config")
var engineOnlyFlag = flag.Bool("engine_only", false, "only state")
var outDirname = flag.String("out", "./tmp", "where to write the files to")

func main() {
	flag.Parse()

	fmt.Println(*engineOnlyFlag)
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
	if err := ioutil.WriteFile(filepath.Join(*outDirname, outFile), code, 0644); err != nil {
		panic(err)
	}

	if err := generateMarshallers(); err != nil {
		panic(err)
	}

	fmt.Println(getstartedfactory.WriteGetStarted(c.State, c.Actions))
}
