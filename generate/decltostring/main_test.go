package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

const outputFilePath string = "testdata/actual_output/output.go"

func TestGenerate(t *testing.T) {

	expected, err := ioutil.ReadFile("testdata/expected_output/golden.go")
	if err != nil {
		t.Fatalf("Error loading golden file: %s", err)
	}

	err = exec.Command(
		"go", "run", ".",
		"-input", "testdata/input/",
		"-output", outputFilePath,
		"-package", "package_name",
		"-prefix", "golden_",
		"-exclude", "fax",
	).Run()

	if err != nil {
		t.Fatalf("Error executing command: %s", err)
	}

	actual, err := ioutil.ReadFile(outputFilePath)
	if err != nil {
		t.Fatalf("Error loading actual output file: %s", err)
	}

	dmp := diffmatchpatch.New()
	fmt.Println("|" + string(actual))
	fmt.Println("|" + string(expected))
	diffs := dmp.DiffMain(string(actual), string(expected), false)

	dmp.DiffPrettyText(diffs)

	if string(expected) != string(actual) {
		t.Errorf(dmp.DiffPrettyText(diffs))
	}
}
