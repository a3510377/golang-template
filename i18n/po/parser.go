package po

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func parse(fileName string) *Message {
	fSet := token.NewFileSet()
	node, _ := parser.ParseFile(fSet, fileName, nil, parser.AllErrors)

	var result *Message

	ast.Inspect(node, func(node ast.Node) bool {
		funcCall, ok := node.(*ast.CallExpr)
		if !ok {
			return true
		}

		funName := ""

		switch n := funcCall.Fun.(type) {
		case *ast.Ident:
			funName = n.Name
		case *ast.SelectorExpr:
			if fun, ok := n.X.(*ast.Ident); ok {
				funName = fun.Name + "." + n.Sel.Name
			}
		}

		if funName != "i18n.Tr" && funName != "tr" {
			return true
		}

		pos := fSet.Position(node.Pos())
		firstArg, ok := (funcCall.Args[0].(*ast.BasicLit))

		if !ok {
			return true
		}

		msg, _ := strconv.Unquote(firstArg.Value)

		result = &Message{
			FileLine: pos.String(),
			Value:    msg,
		}

		return true
	})

	return result
}
