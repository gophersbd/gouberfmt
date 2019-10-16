// Package mutexpointer defines an Analyzer that checks for
// uses of Pointers to sync.Mutex.
// https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid
package mutexpointer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for uses of pointers of sync.Mutex

This checker reports usages of pointers of sync.Mutex in the form of
var x *sync.Mutex. The zero-value of sync.Mutex and sync.RWMutex is valid,
so you almost never need a pointer to a mutex..
`

var Analyzer = &analysis.Analyzer{
	Name: "mutexpointer",
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

type mutexPointerInspector struct {
	pass        *analysis.Pass
	collections []func(node ast.Node) bool
}

func newInspector(pass *analysis.Pass) *mutexPointerInspector {
	inspectors := &mutexPointerInspector{
		pass: pass,
	}
	inspectors.collections = []func(ast.Node) bool{
		inspectors.inspectVariableDeclarations,
		inspectors.inspectInlineVariableDeclarations,
		inspectors.inspectTypeDeclarations,
	}
	return inspectors
}

func (i *mutexPointerInspector) inspectVariableDeclarations(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok {
		return true
	}

	if decl.Tok != token.VAR {
		return true
	}

	ret := true
	for _, spec := range decl.Specs {
		varSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			return true
		}

		pointer, ok := varSpec.Type.(*ast.StarExpr)
		if !ok {
			return true
		}

		if check(pointer.X) {
			return true
		}

		for _, varName := range varSpec.Names {
			i.pass.Reportf(varName.Pos(), "mutex-pointer: var %s uses pointer to mutex", varName.Name)
		}

		ret = false
	}
	return ret
}

func (i *mutexPointerInspector) inspectInlineVariableDeclarations(node ast.Node) bool {
	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return true
	}

	if assignStmt.Tok != token.DEFINE {
		return true
	}

	ret := true
	for index, stmt := range assignStmt.Rhs {
		expr, ok := stmt.(*ast.CallExpr)
		if !ok {
			continue
		}

		caller, ok := expr.Fun.(*ast.Ident)
		if !ok {
			continue
		}

		if caller.Name != "new" {
			continue
		}

		if len(expr.Args) != 1 {
			continue
		}

		arg := expr.Args[0]

		found := check(arg)
		if found {
			continue
		}

		lhsExpr := assignStmt.Lhs[index]

		varName, ok := lhsExpr.(*ast.Ident)
		if !ok {
			continue
		}
		i.pass.Reportf(expr.Pos(), "mutex-pointer: var %s uses pointer to mutex", varName)

		ret = true
	}

	return ret
}

func check(node ast.Expr) bool {
	obj, ok := node.(*ast.SelectorExpr)
	if !ok {
		return true
	}

	ident, ok := obj.X.(*ast.Ident)
	if !ok {
		return true
	}

	if ident.Name != "sync" {
		return true
	}

	if obj.Sel.Name != "Mutex" && obj.Sel.Name != "RWMutex" {
		return true
	}

	return false
}

func (i *mutexPointerInspector) inspectTypeDeclarations(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if !ok {
		return true
	}

	if decl.Tok != token.TYPE {
		return true
	}

	ret := true
	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			return true
		}

		pointer, ok := typeSpec.Type.(*ast.StarExpr)
		if !ok {
			return true
		}

		if check(pointer.X) {
			return true
		}

		i.pass.Reportf(typeSpec.Pos(), "mutex-pointer: type %s points to pointer of mutex", typeSpec.Name)

		ret = false
	}

	return ret
}
