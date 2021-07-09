package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func inspect() {
	fs := http.FileServer(http.Dir(staticServePath()))
	http.Handle("/", fs)

	fmt.Println("inspector running...")
	fmt.Println("open your browser on http://localhost:" + *portFlag)
	fmt.Println("use url query parameter if the server you want to expect ist not listening on 3494")
	fmt.Println("eg. http://localhost:" + *portFlag + "?port=8080")

	err := http.ListenAndServe(":"+*portFlag, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func staticServePath() string {
	if *devModeFlag {
		return "inspector/build"
	}
	return gopath() + "/src/github.com/Java-Jonas/bar-cli/inspector/build"
}

func gopath() string {
	p := os.Getenv("GOPATH")

	if len(p) == 0 {
		panic("could not find $GOPATH")
	}

	return p
}
