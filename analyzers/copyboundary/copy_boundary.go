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

func (c *copyBoundaryInspector) inspectAssignStatement(node ast.Node) bool {
	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return true
	}

	for i := 0; i < len(assignStmt.Rhs); i++ {
		lhs := assignStmt.Lhs[i]

		ident, ok := lhs.(*ast.Ident)

		if !ok || ident.Obj == nil || ident.Name == "_" {
			continue
		}

		typ := isSliceOrMap(assignStmt.Rhs[i])
		if typ == "" {
			continue
		}

		c.pass.Reportf(lhs.Pos(), "copy-boundary: copies a %s directly", typ)
	}

	return false
}

func isSliceOrMap(rootExpr ast.Expr) (typ string) {
	obj := GetObj(rootExpr)
	if obj == nil {
		return
	}
	decl := obj.Decl
	expr := GetExpr(decl)

	if expr == nil {
		assignStmt, ok := decl.(*ast.AssignStmt)
		if !ok {
			return
		}

		for i := 0; i < len(assignStmt.Rhs); i++ {
			childObj := GetObj(assignStmt.Lhs[i])
			if childObj == nil {
				continue
			}

			if obj.Name == childObj.Name {
				expr = GetExpr(assignStmt.Rhs[i])
				break
			}
		}
	}

	switch expr.(type) {
	case *ast.ArrayType:
		typ = "slice"
	case *ast.MapType:
		typ = "map"
	}

	return
}

func GetObj(expr ast.Expr) *ast.Object {
	ident, ok := expr.(*ast.Ident)
	if !ok {
		return nil
	}

	return ident.Obj
}

func GetExpr(decl interface{}) ast.Expr {
	switch decl := decl.(type) {
	case *ast.ValueSpec:
		return decl.Type
	case *ast.Field:
		return decl.Type
	case *ast.CompositeLit:
		return decl.Type
	}

	return nil
}
