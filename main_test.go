package statefactory

import (
	"bytes"
	"strings"
	"unicode"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func simplifyIfWhitespace(ch rune) rune {
	if ch == '\n' {
		return ch
	}
	if unicode.IsSpace(ch) {
		return ' '
	}
	return ch
}

func normalizeWhitespace(_str string) string {
	str := strings.TrimSpace(_str)
	var b strings.Builder
	b.Grow(len(str))

	var lastWrittenRune rune = '1'

	for _, _ch := range str {
		ch := simplifyIfWhitespace(_ch)
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
			lastWrittenRune = ch
		} else {
			if lastWrittenRune != ch {
				b.WriteRune(ch)
				lastWrittenRune = ch
			}
		}
	}

	return b.String()
}

// creates diff and makes whitespace visible
func diff(actual, expected string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(actual, expected, true)

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

DIFF:
` + dmp.DiffPrettyText(diffs) + `

__________________________________

WANT:
` + expected + `

GOT:
` + actual
}

func newSimpleASTExample() simpleAST {
	data := map[interface{}]interface{}{
		"player": map[interface{}]interface{}{
			"items":     "[]item",
			"gearScore": "gearScore",
			"position":  "position",
		},
		"zone": map[interface{}]interface{}{
			"items":   "[]zoneItem",
			"players": "[]player",
		},
		"zoneItem": map[interface{}]interface{}{
			"position": "position",
			"item":     "item",
		},
		"position": map[interface{}]interface{}{
			"x": "float64",
			"y": "float64",
		},
		"item": map[interface{}]interface{}{
			"gearScore": "gearScore",
		},
		"gearScore": map[interface{}]interface{}{
			"level": "int",
			"score": "int",
		},
	}

	// TODO: make prettier
	simpleAST := buildRudimentarySimpleAST(data)
	simpleAST.fillInReferences().fillInParentalInfo()

	return simpleAST
}
