package actionsfactory

import (
	"bytes"
	"sort"
	"strings"
	"text/template"
)

type actionsFactory struct {
	ast *actionsConfigAST
	buf *bytes.Buffer
}

func newActionsFactory(ast *actionsConfigAST) *actionsFactory {
	return &actionsFactory{
		ast: ast,
		buf: &bytes.Buffer{},
	}
}

func (s *actionsFactory) writtenSourceCode() []byte {
	return s.buf.Bytes()
}

// indexOfAction is a helper function for finding the index of a given action
// within the actionsConfig. Since golang's templates loop through maps (actionsConfigAST is a map)
// in alphabetical order, this will give a deterministic output within the templating frame
func indexOfParam(params map[string]actionParameter, currentParam actionParameter) int {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var indexOf int
	for i, key := range keys {
		if key == currentParam.Name {
			indexOf = i
		}
	}
	return indexOf
}

func newTemplateFrom(name, templateString string) *template.Template {
	return template.Must(template.
		New(name).
		Funcs(template.FuncMap{
			"toTitleCase": strings.Title,
			// does not write given string at certain index of param (determined by alphabetical order of paramsConfigAST)
			"doNotWriteOnIndex": func(params map[string]actionParameter, currentParam actionParameter, requiredIndex int, toWrite string) string {
				currentIndex := indexOfParam(params, currentParam)
				if requiredIndex < 0 {
					if currentIndex == len(params)+requiredIndex {
						return ""
					}
				} else {

					if currentIndex == requiredIndex {
						return ""
					}
				}
				return toWrite
			},
			// does only write given string at certain index of param (determined by alphabetical order of paramsConfigAST)
			"writeOnIndex": func(params map[string]actionParameter, currentParam actionParameter, requiredIndex int, toWrite string) string {
				currentIndex := indexOfParam(params, currentParam)
				if requiredIndex < 0 {
					if currentIndex == len(params)+requiredIndex {
						return toWrite
					}
				} else {

					if currentIndex == requiredIndex {
						return toWrite
					}
				}
				return ""
			},
		}).
		Delims("<(", ")>").
		Parse(templateString))
}
