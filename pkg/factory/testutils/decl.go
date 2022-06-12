package testutils

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	_ast "github.com/jobergner/backent-cli/pkg/ast"
	"github.com/jobergner/backent-cli/pkg/factory/configs"
)

func evalDeclName(decl ast.Decl) string {
	if isFuncDecl(decl) {
		declName := getFuncName(decl.(*ast.FuncDecl)) + "_func"
		return ensureLowerCase(declName)
	}
	if isGenDecl(decl) {
		declName := getGenDeclName(decl.(*ast.GenDecl)) + "_type"
		return ensureLowerCase(declName)
	}
	panic("unknown decl name")
}

func isFuncDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.FuncDecl); ok {
		return true
	}
	return false
}

func isGenDecl(decl ast.Decl) bool {
	if _, ok := decl.(*ast.GenDecl); ok {
		return true
	}
	return false
}

func getFuncName(decl *ast.FuncDecl) string {
	if decl.Recv == nil {
		return decl.Name.Name
	}
	receiverName := findMethodReceiverIdent(decl.Recv.List[0].Type)
	return decl.Name.Name + "_" + receiverName.Name
}

func findMethodReceiverIdent(expr ast.Expr) *ast.Ident {
	if identType, ok := expr.(*ast.Ident); ok {
		return identType
	}
	if arrayType, ok := expr.(*ast.ArrayType); ok {
		return findMethodReceiverIdent(arrayType.Elt)
	}
	return findMethodReceiverIdent(expr.(*ast.StarExpr).X)
}

func getGenDeclName(decl *ast.GenDecl) string {
	if typeSpec, ok := decl.Specs[0].(*ast.TypeSpec); ok {
		return typeSpec.Name.Name
	}
	return decl.Specs[0].(*ast.ValueSpec).Names[0].Name
}

func ensureLowerCase(s string) string {
	if string(s[0]) == strings.ToUpper(string(s[0])) {
		return "_" + s
	}
	return s
}

func parseDecls(code string) []ast.Decl {
	ast, err := parser.ParseFile(token.NewFileSet(), "", PackageClause+code, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	return ast.Decls
}

func stringifyDecl(decl ast.Decl) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), decl)
	return buf.String()
}

func isImportDecl(decl ast.Decl) bool {
	if genDecl, ok := decl.(*ast.GenDecl); ok {
		if genDecl.Tok == token.IMPORT {
			return true
		}
	}
	return false
}

func parseDeclsToString(code string) []string {
	decls := parseDecls(code)

	var stringifiedDecls []string
	for _, decl := range decls {
		stringifiedDecls = append(stringifiedDecls, stringifyDecl(decl))
	}

	return stringifiedDecls
}

func FindUnmatchedDecls(code string, collection map[string]string) ([]string, bool) {
	declMatcher := make(map[string]struct{})
	for k := range collection {
		declMatcher[k] = struct{}{}
	}

	actualDecls := parseDeclsToString(code)

	for declName, expectedDecl := range collection {
		for _, actualDecl := range actualDecls {
			if actualDecl != expectedDecl {
				continue
			}

			delete(declMatcher, declName)
		}
	}

	var unmatchedDecls []string
	for k := range declMatcher {
		unmatchedDecls = append(unmatchedDecls, k)
	}

	return unmatchedDecls, len(declMatcher) > 0
}

func FindRedundantDecls(code string, collection map[string]string) ([]string, bool) {
	actualDecls := parseDecls(code)

	declMatcher := make(map[string]struct{})
	// TODO: this requires evalDeclName of this package to equal declToString's evalDeclName. it's time to merge repos
	for _, d := range actualDecls {
		declMatcher[evalDeclName(d)] = struct{}{}
	}

	for declName := range collection {
		delete(declMatcher, declName)
	}

	var redundantDecls []string
	for k := range declMatcher {
		redundantDecls = append(redundantDecls, k)
	}

	return redundantDecls, len(declMatcher) > 0
}

func NewSimpleASTExample() *_ast.AST {
	simpleAST := _ast.Parse(configs.StateConfig, map[interface{}]interface{}{}, map[interface{}]interface{}{})
	return simpleAST
}
