// Package copyboundary defines an Analyzer that checks for
// uses of copy of slice & map.
// https://github.com/uber-go/guide/blob/master/style.md#copy-slices-and-maps-at-boundaries
package copyboundary

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for uses of copy of slice & map

This checker reports usages of slice & map that are directly copied. 
Slices and maps contain pointers to the underlying data so be wary 
of scenarios when they need to be copied..
`

var Analyzer = &analysis.Analyzer{
	Name: "copyboundary",
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

type copyBoundaryInspector struct {
	pass        *analysis.Pass
	collections []func(node ast.Node) bool
}

func newInspector(pass *analysis.Pass) *copyBoundaryInspector {
	inspectors := &copyBoundaryInspector{
		pass: pass,
	}
	inspectors.collections = []func(ast.Node) bool{
		inspectors.inspectAssignStatement,
	}
	return inspectors
}

func (i *copyBoundaryInspector) inspectAssignStatement(node ast.Node) bool {
	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return true
	}

	typ := checkType(assignStmt.Lhs)
	if typ == "" {
		return true
	}

	typ = checkType(assignStmt.Rhs)
	if typ == "" {
		return true
	}

	i.pass.Reportf(assignStmt.Pos(), "copy-boundary: copies a %s directly", typ)

	return false
}

func checkType(obj []ast.Expr) string {
	if len(obj) > 1 {
		return ""
	}

	x, ok := obj[0].(*ast.Ident)
	if !ok {
		return ""
	}

	var typ ast.Expr

	switch decl := x.Obj.Decl.(type) {
	case *ast.ValueSpec:
		typ = decl.Type
	case *ast.Field:
		typ = decl.Type
	default:
		return ""
	}

	switch typ.(type) {
	case *ast.ArrayType:
		return "slice"
	case *ast.MapType:
		return "map"
	}

	return ""
}
