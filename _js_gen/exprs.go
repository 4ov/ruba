package jsgen

import (
	"fmt"
	"plang/ast"
	"strings"
)

func expression(expr ast.Expression) string {
	switch expr := expr.(type) {
	case ast.AddExpr:
		return addExpr(expr)
	case ast.SubExpr:
		return subExpr(expr)
	case ast.DivExpr:
		return divExpr(expr)
	case ast.MulExpr:
		return mulExpr(expr)
	case ast.IdentExpr:
		return identExpr(expr)
	case ast.NumberExpr:
		return expr.Value
	case ast.StringExpr:
		return stringExpr(expr)
	case ast.DotExpr:
		return dotExpr(expr)
	case ast.ArrayExpr:
		return arrayExpr(expr)
	case ast.ObjectExpr:
		return objectExpr(expr)
	case ast.EqualExpr:
		return equalExpr(expr)
	case ast.InternalConstExpr:
		return internalConstExpr(expr)
	}
	return ""
}

func dotExpr(expr ast.DotExpr) string {
	// fmt.Println(expr.Parent)
	// return fmt.Sprint("xd")
	return ""
}

func addExpr(expr ast.AddExpr) string {
	right := expression(expr.Right)
	left := expression(expr.Left)
	return fmt.Sprintf("%v + %v", left, right)
}

func equalExpr(expr ast.EqualExpr) string {
	right := expression(expr.Right)
	left := expression(expr.Left)
	return fmt.Sprintf("%v == %v", left, right)
}

func subExpr(expr ast.SubExpr) string {
	right := expression(expr.Right)
	left := expression(expr.Left)
	return fmt.Sprintf("%v - %v", left, right)
}

func mulExpr(expr ast.MulExpr) string {
	right := expression(expr.Right)
	left := expression(expr.Left)
	return fmt.Sprintf("%v * %v", left, right)
}

func divExpr(expr ast.DivExpr) string {
	right := expression(expr.Right)
	left := expression(expr.Left)
	return fmt.Sprintf("%v / %v", left, right)
}

func identExpr(expr ast.IdentExpr) string {
	return expr.Value
}

func internalConstExpr(expr ast.InternalConstExpr) string {
	return fmt.Sprintf("CONSTS.%v", expr.Value)
}

func stringExpr(expr ast.StringExpr) string {
	return fmt.Sprintf(`'%v'`, expr.Value)
}

func arrayExpr(expr ast.ArrayExpr) string {
	exps := []string{}
	for _, exp := range expr.Elements {
		exps = append(exps, expression(exp))
	}
	return fmt.Sprintf("[%v]", strings.Join(exps, ","))
}

func objectExpr(expr ast.ObjectExpr) string {
	result := "{"

	for key, value := range expr.Elements {
		switch key := key.(type) {
		case string:
			result += fmt.Sprintf("%v : %v,", key, expression(value))
		default:
			result += fmt.Sprintf("[%v]: %v,", expression(key), expression(value))
		}
		// fmt.Println(key, value)
	}

	result += "}"
	return result
}
