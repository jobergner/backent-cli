package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"sort"
	"strings"
	"unicode"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

var (
	PackageName   = "main"
	PackageClause = fmt.Sprintf("package %s\n", PackageName)
)

func caseInsensitiveSort(keys []string) func(i, j int) bool {
	return func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	}
}

func Diff(actual, expected string) (string, bool) {
	a := parseDecls(actual)
	b := parseDecls(expected)

	var areDifferent bool

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
			areDifferent = true
			buf.WriteString(fmt.Sprintf("expected to find '%s' but did not\n\n", name))
		}
	}
	for name := range actualDelcs {
		if _, ok := expectedDelcs[name]; !ok {
			areDifferent = true
			buf.WriteString(fmt.Sprintf("found '%s' but should not have\n\n", name))
		}
	}

	var names []string
	for name := range expectedDelcs {
		names = append(names, name)
	}

	sort.Slice(names, caseInsensitiveSort(names))

	for _, name := range names {
		got := actualDelcs[name]
		want := expectedDelcs[name]
		if got != want {
			areDifferent = true
			buf.WriteString(diffDecl(got, want))
		} else {
			buf.WriteString("\n____\n\n CORRECT:\n" + want)
		}
	}

	return buf.String(), areDifferent
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

// FormatUnpackagedCode returns formatted code
func FormatUnpackagedCode(code string) string {
	return FormatCode(PackageClause + code)
}

// FormatUnpackagedCode returns formatted code without a package name
func FormatCode(code string) string {
	ast, err := parser.ParseFile(token.NewFileSet(), "", code, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = format.Node(&buf, token.NewFileSet(), ast)
	if err != nil {
		panic(err)
	}

	return strings.TrimPrefix(buf.String(), PackageClause)
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
			panic(err)
		}
		return diffString
	}

	return ""
}
