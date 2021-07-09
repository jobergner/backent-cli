//go:generate bash ./generate.sh
package main

import (
	"flag"
	"fmt"
	"os"
)

const outFile = "state.go"

var configNameFlag = flag.String("config", "./example.config.json", "path of config")
var engineOnlyFlag = flag.Bool("engine_only", false, "only state")
var outDirName = flag.String("out", "./tmp", "where to write the files to")
var exampleFlag = flag.Bool("example", false, "when enabled starts example")
var devModeFlag = flag.Bool("dev", false, "start in dev mode")
var portFlag = flag.String("port", "3100", "start in dev mode")

func main() {
	flag.Parse()

	var args []string
	for _, arg := range os.Args {
		if arg[0:1] != "-" && arg[0:1] != "--" {
			args = append(args, arg)
		}
	}

	if len(args) < 2 {
		fmt.Println("available commands: `generate`, `inspect`")
		os.Exit(1)
	}

	switch args[1] {
	case "inspect":
		inspect()
	case "generate":
		generate()
	default:
		panic("unknown command: " + args[1])
	}
}
