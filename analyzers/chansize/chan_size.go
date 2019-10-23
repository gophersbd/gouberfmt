// Package chansize defines an Analyzer that checks for
// uses of channel with size greater than 1
// https://github.com/uber-go/guide/blob/master/style.md#channel-size-is-one-or-none
package chansize

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for uses of channel with size greater than 1

Channels should usually have a size of one or be unbuffered. 
By default, channels are unbuffered and have a size of zero. 
Any other size must be subject to a high level of scrutiny.
`

var Analyzer = &analysis.Analyzer{
	Name: "chansize",
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

type chanSizeInspector struct {
	pass        *analysis.Pass
	collections []func(node ast.Node) bool
}

func newInspector(pass *analysis.Pass) *chanSizeInspector {
	inspectors := &chanSizeInspector{
		pass: pass,
	}
	inspectors.collections = []func(ast.Node) bool{
		inspectors.inspectAssignStatement,
	}
	return inspectors
}

func (c *chanSizeInspector) inspectAssignStatement(node ast.Node) bool {
	assignStmt, ok := node.(*ast.AssignStmt)
	if !ok {
		return true
	}

	for _, rhs := range assignStmt.Rhs {
		callExpr, ok := rhs.(*ast.CallExpr)
		if !ok {
			continue
		}

		ident, ok := callExpr.Fun.(*ast.Ident)
		if !ok {
			continue
		}

		if ident.Name != "make" {
			continue
		}

		if len(callExpr.Args) == 1 {
			continue
		}

		_, ok = callExpr.Args[0].(*ast.ChanType)
		if !ok {
			continue
		}

		basicLit, ok := callExpr.Args[1].(*ast.BasicLit)
		if !ok {
			continue
		}

		if basicLit.Value == "1" {
			continue
		}

		c.pass.Reportf(rhs.Pos(), "chan-size: channel size should be one or none")
	}

	return false
}
