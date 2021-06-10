package main

import (
	"bar-cli/enginefactory"
	"bar-cli/examples"
	"bar-cli/factoryutils"
	"bar-cli/serverfactory"
	"bytes"
	"fmt"
	"io/ioutil"
)

var relevantFiles = [5]string{
	"./serverfactory/server_example/server/client.go",
	"./serverfactory/server_example/server/connect.go",
	"./serverfactory/server_example/server/main.go",
	"./serverfactory/server_example/server/message.go",
	"./serverfactory/server_example/server/room.go",
}

// TODO handle imports
func writeRelevantFiles(buf *bytes.Buffer) {
	for _, relevantFile := range relevantFiles {
		buf.WriteString("\n")
		dat, err := ioutil.ReadFile(relevantFile)
		if err != nil {
			panic(err)
		}
		buf.WriteString(factoryutils.TrimPackageName(string(dat)))
	}
}

func main() {
	buf := bytes.NewBufferString("package main\n")
	enginefactory.WriteEngine(buf, examples.StateConfig)
	serverfactory.WriteServer(buf, examples.StateConfig, examples.ActionsConfig)
	writeRelevantFiles(buf)
	fmt.Println(buf.String())
}
