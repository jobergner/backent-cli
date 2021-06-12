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

	"encoding/json"
)

const outDir = "./tmp"
const outFile = "state.go"

type config struct {
	State   map[string]interface{} `json:"state"`
	Actions map[string]interface{} `json:"actions"`
}

var configNameFlag = flag.String("config", "./barcli.config.json", "path of config")

func atob(a map[string]interface{}) map[interface{}]interface{} {
	b := make(map[interface{}]interface{})
	for k, v := range a {
		x := make(map[interface{}]interface{})
		// TODO a LOT!
		// what if a string and not map[string]interface{}
		for k_, v_ := range v.(map[string]interface{}) {
			x[k_] = v_
		}
		b[k] = x
	}
	return b
}

func main() {
	configFile, err := ioutil.ReadFile(*configNameFlag)
	if err != nil {
		panic(err)
	}
	c := config{}
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		panic(err)
	}
	fmt.Println(atob(c.State))
	if errs := validator.ValidateStateConfig(atob(c.State)); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("errors")
	}
	if errs := validator.ValidateActionsConfig(atob(c.State), atob(c.Actions)); len(errs) != 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		panic("errors")
	}
	buf := bytes.NewBufferString("package state\n")
	writeCombinedImport(buf)
	writeImportedFiles(buf)
	enginefactory.WriteEngine(buf, atob(c.State))
	serverfactory.WriteServer(buf, atob(c.State), atob(c.Actions))
	if err := ioutil.WriteFile(filepath.Join(outDir, outFile), buf.Bytes(), 0644); err != nil {
		panic(err)
	}
	cmd := exec.Command("easyjson", "-all", "-omit_empty", filepath.Join(outDir, outFile))
	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
	fmt.Println(getstartedfactory.WriteGetStarted(atob(c.State), examples.ActionsConfig))
}
