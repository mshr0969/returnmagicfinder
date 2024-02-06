package returnmagicfinder

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "returnmagicfinder is a tool to identify magic numbers in return statement."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "returnmagicfinder",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		switch x := n.(type) {
		case *ast.ReturnStmt:
			reported := false
			for _, expr := range x.Results {
				if reported {
					break
				}
				ast.Inspect(expr, func(n ast.Node) bool {
					if lit, ok := n.(*ast.BasicLit); ok && (lit.Kind == token.INT || lit.Kind == token.FLOAT) && !reported {
						pass.Reportf(lit.Pos(), "magic number used in return statement")
						reported = true
					}
					return true
				})
			}
		}
	})

	return nil, nil
}
