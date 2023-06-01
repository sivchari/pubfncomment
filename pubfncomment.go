package pubfncomment

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "pubfncomment is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "pubfncomment",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return
		}
		if !unicode.IsUpper(rune(fn.Name.Name[0])) {
			return
		}
		if fn.Doc == nil {
			pass.Reportf(fn.Pos(), "public function (%s) should have comment", fn.Name.Name)
		}
	})
	return nil, nil
}
