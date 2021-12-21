package cgen

import (
	"fmt"
	"plang/ast"
	"reflect"
)

func expression(expr ast.Expression) string {
	switch expr := expr.(type) {
	case ast.IdentExpr:
		return expr.Value
	case ast.IntExpr:
		return expr.Value
	case ast.FloatExpr:
		return expr.Value
	case ast.EqualExpr:
		return fmt.Sprintf("%v == %v", expression(expr.Left), expression(expr.Right))
	case string:
		return expr

	default:
		fmt.Println("!!", reflect.TypeOf(expr))
	}
	return ""
}
