// Package interfacepointer defines an Analyzer that checks for
// uses of Pointers to Interface.
// https://github.com/uber-go/guide/blob/master/style.md#pointers-to-interfaces
package interfacepointer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for uses of pointers of interface

This checker reports usages of pointers of interface in the form of
var x *interface{}. Usually passing interfaces as values is recommended — 
the underlying data can still be a pointer.
`

var Analyzer = &analysis.Analyzer{
	Name: "interfacepointer",
	Doc:  Doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := newInspector(pass)
	for _, file := range pass.Files {
		for _, inspect := range inspector.collections {
			ast.Inspect(file, inspect)
		}
	}
	return nil, nil
}

type interfacePointerInspector struct {
	pass        *analysis.Pass
	collections []func(node ast.Node) bool
}

func newInspector(pass *analysis.Pass) *interfacePointerInspector {
	inspectors := &interfacePointerInspector{
		pass: pass,
	}
	inspectors.collections = []func(ast.Node) bool{
		inspectors.inspectVariableDeclarations,
		inspectors.inspectFunctionDeclarations,
	}
	return inspectors
}

func (i *interfacePointerInspector) inspectVariableDeclarations(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok {
		return true
	}

	if decl.Tok != token.VAR {
		return true
	}

	var ret = true
	for _, spec := range decl.Specs {
		varSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			return true
		}

		pointer, ok := varSpec.Type.(*ast.StarExpr)
		if !ok {
			return true
		}

		_, ok = pointer.X.(*ast.InterfaceType)
		if !ok {
			return true
		}

		ret = false
		for _, varName := range varSpec.Names {
			i.pass.Reportf(varName.Pos(), "interface-pointer: var %s uses pointer to interface", varName.Name)
		}
	}
	return ret
}

// TODO Implement
func (i *interfacePointerInspector) inspectFunctionDeclarations(node ast.Node) bool {
	return true
}
