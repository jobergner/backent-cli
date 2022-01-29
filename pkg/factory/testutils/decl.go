package testutils

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
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
	packageClause := "package main\n"

	ast, err := parser.ParseFile(token.NewFileSet(), "", packageClause+code, parser.AllErrors)
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
