package inter

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/4ov/ruba/ast"
)

func expression(expr ast.Expression, env *Env) IType {
	switch expr := expr.(type) {
	case ast.IntExpr:
		return intExpr(expr)
	case ast.MathExpr:
		return mathExpr(expr, env)
	case ast.IdentExpr:
		return identExpr(expr, env)
	case ast.DotExpr:
		return dotExpr(expr, env)
	case ast.IndexExpr:
		return indexExpr(expr, env)
	case ast.ObjectExpr:
		return objectExpr(expr, env)
	case ast.StringExpr:
		return stringExpr(expr, env)
	case ast.ArrayExpr:
		return arrayExpr(expr, env)
	case ast.ArgExpr:
		return argExpr(expr, env)
	case string:
		return NewString(expr)
	case ast.FnCallExpr:
		return fnCallExpr(expr, env)
	default:
		panic(fmt.Sprintf("unsupported expression %v", reflect.TypeOf(expr).Name()))
	}
}

func intExpr(expr ast.IntExpr) IType {
	i, err := strconv.Atoi(expr.Value)
	if err != nil {
		panic(err)
	}
	return NewInt(i)
}

func mathExpr(expr ast.MathExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	switch expr.Operator {
	case "+":
		return Add(l, r)
	case "*":
		return Mul(l, r)
	default:
		panic("not implemented")
	}

}

func identExpr(expr ast.IdentExpr, env *Env) IType {
	return env.Get(expr.Value)
}

func dotExpr(expr ast.DotExpr, env *Env) IType {
	fmt.Println(expr)
	r := expression(expr.Parent, env).Access(expression(expr.Child, env))
	return r
}

func indexExpr(expr ast.IndexExpr, env *Env) IType {
	r := expression(expr.Parent, env).Access(expression(expr.Child, env))
	return r
}

func stringExpr(expr ast.StringExpr, env *Env) IType {

	return NewString(expr.Value)
}

func objectExpr(expr ast.ObjectExpr, env *Env) IType {
	obj := NewObject(ObjectValue{})
	for key, value := range expr.Elements {
		if key := expression(key, env); key.TypeOf() == STRING {
			obj.Set(key.String(), expression(value, env))
		} else {
			panic("key must be string")
		}

	}
	return obj
}

func arrayExpr(expr ast.ArrayExpr, env *Env) IType {
	arr := NewArray([]IType{})
	for _, value := range expr.Elements {
		arr.Value = append(arr.Value, expression(value, env))
	}
	return arr
}

func argExpr(expr ast.ArgExpr, env *Env) IType {
	if !expr.Spreadable {
		return expression(expr.Value, env)

	} else {
		switch value := expression(expr.Value, env).(type) {
		case *Array:
			return value
		case *String:
			result := NewArray([]IType{})
			for _, char := range value.Value {
				result.Value = append(result.Value, NewString(string(char)))
			}
			return result
		default:
			panic("spreadable argument must be array or string")
		}
	}
}

func fnCallExpr(stmt ast.FnCallExpr, env *Env) IType {
	var target IType
	//I GUESS THIS IS A BUG

	switch raw := stmt.Target.(type) {
	case string:
		target = env.Get(raw)
	default:
		target = expression(raw, env)
	}

	//nilAsNull will turn go's nil into usable value from ruba
	switch target := target.(type) {
	case *NativeFn:
		return nilAsNull(CallNativeFn(target, stmt.Args, env))
	case *Fn:
		return nilAsNull(CallFn(target, stmt.Args, env))
	default:
		panic(fmt.Sprintf("%s is not a function", target))
	}
}
