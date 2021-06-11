package main

import (
	"bar-cli/enginefactory"
	"bar-cli/examples"
	"bar-cli/getstartedfactory"
	"bar-cli/serverfactory"
	"bytes"
	"flag"
	"fmt"
	"github.com/Java-Jonas/validator"
	"io/ioutil"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const outDir = "./tmp"
const outFile = "state.go"

type config struct {
	State   map[interface{}]interface{} `yaml:"state"`
	Actions map[interface{}]interface{} `yaml:"actions"`
}

var configNameFlag = flag.String("config", "./barcli.config.yaml", "path of config")

func main() {
	configFile, err := ioutil.ReadFile(*configNameFlag)
	if err != nil {
		panic(err)
	}
	c := config{}
	err = yaml.Unmarshal([]byte(configFile), &c)
	if err != nil {
		panic(err)
	}
	if errs := validator.ValidateStateConfig(c.State); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("errors")
	}
	if errs := validator.ValidateActionsConfig(c.State, c.Actions); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("errors")
	}
	fmt.Printf("%+v\n", c)
	buf := bytes.NewBufferString("package state\n")
	writeCombinedImport(buf)
	writeImportedFiles(buf)
	enginefactory.WriteEngine(buf, c.State)
	serverfactory.WriteServer(buf, c.State, c.Actions)
	if err := ioutil.WriteFile(filepath.Join(outDir, outFile), buf.Bytes(), 0644); err != nil {
		panic(err)
	}
	cmd := exec.Command("easyjson", "-all", "-omit_empty", filepath.Join(outDir, outFile))
	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
	fmt.Println(getstartedfactory.WriteGetStarted(examples.StateConfig, examples.ActionsConfig))
}
