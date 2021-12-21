package checker

import (
	"fmt"
	"plang/ast"
)

func expression(expr ast.Expression, memory *Memory) int {
	switch expr := expr.(type) {
	case ast.AddExpr:
		return addExpr(expr, memory)
	case ast.SubExpr:
		return subExpr(expr, memory)
	case ast.DivExpr:
		return divExpr(expr, memory)
	case ast.MulExpr:
		return mulExpr(expr, memory)
	case ast.IdentExpr:
		return identExpr(expr, memory)
	case ast.NumberExpr:
		return ast.NUMBERT
	case ast.StringExpr:
		return ast.STRINGT
	case ast.DotExpr:
		return dotExpr(expr, memory)
	default:
		// fmt.Println(expr, "xx")
		return -1
	}

}

func dotExpr(expr ast.DotExpr, memory *Memory) int {
	// fmt.Println(expr.Parent)
	return ast.NUMBERT
}

func addExpr(expr ast.AddExpr, memory *Memory) int {
	right := expression(expr.Right, memory)
	left := expression(expr.Left, memory)
	if right == left {
		return ast.NUMBERT
	} else {
		panic("")
	}

}

func subExpr(expr ast.SubExpr, memory *Memory) int {
	right := expression(expr.Right, memory)
	left := expression(expr.Left, memory)
	if right == ast.NUMBERT && left == ast.NUMBERT {
		return ast.NUMBERT
	} else {
		panic("")
	}
}

func mulExpr(expr ast.MulExpr, memory *Memory) int {
	right := expression(expr.Right, memory)
	left := expression(expr.Left, memory)
	if right == ast.NUMBERT && left == ast.NUMBERT {
		return ast.NUMBERT
	} else {
		panic("")
	}
}

func divExpr(expr ast.DivExpr, memory *Memory) int {
	right := expression(expr.Right, memory)
	left := expression(expr.Left, memory)
	if right == ast.NUMBERT && left == ast.NUMBERT {
		return ast.NUMBERT
	} else {
		panic("")
	}
}

func identExpr(expr ast.IdentExpr, memory *Memory) int {
	if ival, ok := memory.Idents[expr.Value]; ok {
		return ival
	} else {
		panic(fmt.Sprintf("%v is not defined", expr.Value))
	}
}

func stringExpr(expr ast.StringExpr, memory *Memory) int {
	return ast.STRINGT
}
