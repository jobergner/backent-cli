package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

func evalDeclName(decl ast.Decl, containingFileName string) string {
	if isImportDecl(decl) {
		declName := evalImportName(containingFileName) + "_import"
		return ensureLowerCase(declName)
	}
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

func ensureLowerCase(s string) string {
	if string(s[0]) == strings.ToUpper(string(s[0])) {
		return "_" + strings.ToLower(string(s[0])) + string(s[1:len(s)])
	}
	return s
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

// file name "foo.bar" returns "foo_bar"
func evalImportName(containingFileName string) string {
	return strings.Replace(containingFileName, ".", "_", -1)
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
