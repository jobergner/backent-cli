package actionsfactory

import (
	"bytes"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
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

// WriteActionsFrom writes source code for a given ActionsConfig
// moduleName is the name of the module the state is created for
// packageName is the name of the package the state is used in
func WriteActionsFrom(actionsConfigData map[interface{}]interface{}, moduleName string, packageName string) []byte {
	actionsConfigAST := buildActionsConfigAST(actionsConfigData)
	a := newActionsFactory(actionsConfigAST).
		writePackageName(packageName).
		writeImport(moduleName).
		writeStart().
		writeActions()

	err := a.format()
	if err != nil {
		// unexpected error
		fmt.Println(string(a.writtenSourceCode()))
		panic(err)
	}

	return a.writtenSourceCode()
}

func (s *actionsFactory) writeImport(moduleName string) *actionsFactory {
	importDecl := `import (
	"` + moduleName + `/state"
)
`
	s.buf.WriteString(importDecl)
	return s
}

func (s *actionsFactory) writePackageName(packageName string) *actionsFactory {
	s.buf.WriteString("package " + packageName + "\n\n")
	return s
}

func (a *actionsFactory) format() error {
	ast, err := parser.ParseFile(token.NewFileSet(), "", a.buf.String(), parser.AllErrors)
	if err != nil {
		return err
	}

	a.buf.Reset()
	err = format.Node(a.buf, token.NewFileSet(), ast)
	return err
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
