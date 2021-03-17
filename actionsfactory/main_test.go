package actionsfactory

// TODO duplicate!!

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

// writing templates can be irritating if you have to be
// precise with whitespace. this normalization function
// takes away some of that irritation by removing all consecutive
// whitespace of a certain kind (newline or everything else) except
// for one
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
