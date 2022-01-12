package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"unicode"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

func Diff(actual, expected string) string {
	a := parseDecls(actual)
	b := parseDecls(expected)

	actualDelcs := make(map[string]string)
	for _, decl := range a {
		actualDelcs[evalDeclName(decl)] = stringifyDecl(decl)
	}
	expectedDelcs := make(map[string]string)
	for _, decl := range b {
		expectedDelcs[evalDeclName(decl)] = stringifyDecl(decl)
	}

	var buf bytes.Buffer
	for name := range expectedDelcs {
		if _, ok := actualDelcs[name]; !ok {
			buf.WriteString(fmt.Sprintf("expected to find '%s' but did not\n\n", name))
		}
	}
	for name := range actualDelcs {
		if _, ok := expectedDelcs[name]; !ok {
			buf.WriteString(fmt.Sprintf("found '%s' but should not have\n\n", name))
		}
	}

	for name, def := range expectedDelcs {
		got := actualDelcs[name]
		want := def
		if got != want {
			buf.WriteString(diffDecl(got, want))
		}
	}

	return buf.String()
}

// creates diff and makes whitespace visible
func diffDecl(actual, expected string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(actual, expected, false)

	for i, diff := range diffs {
		if diff.Type == diffmatchpatch.DiffDelete || diff.Type == diffmatchpatch.DiffInsert {
			var buf bytes.Buffer
			for _, ch := range diff.Text {
				if unicode.IsSpace(ch) {
					buf.WriteString("~")
				} else {
					buf.WriteRune(ch)
				}
			}
			diff.Text = buf.String()
			diffs[i] = diff
		}
	}

	return `
__________________________________
DIFF:
` + dmp.DiffPrettyText(diffs) + `


WANT:
` + expected + `

GOT:
` + actual
}

func FormatCode(code string) string {
	packageClause := "package main\n"

	ast, err := parser.ParseFile(token.NewFileSet(), "", packageClause+code, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = format.Node(&buf, token.NewFileSet(), ast)
	return strings.TrimPrefix(buf.String(), packageClause)
}

func DiffJSON(actual, expected string) string {
	aString := []byte(actual)
	bString := []byte(expected)

	differ := gojsondiff.New()
	d, err := differ.Compare(aString, bString)
	if err != nil {
		panic(err)
	}

	if d.Modified() {
		var aJson map[string]interface{}
		json.Unmarshal(aString, &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		formatter := formatter.NewAsciiFormatter(aJson, config)
		diffString, err := formatter.Format(d)
		if err != nil {
			// No error can occur
		}
		return diffString
	}

	return ""
}
