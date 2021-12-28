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
	case ast.FloatExpr:
		return floatExpr(expr)
	case ast.BoolExpr:
		return boolExpr(expr)
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
	case ast.EqualExpr:
		return equalExpr(expr, env)
	case ast.NullExpr:
		return nullExpr(expr)
	case ast.NotExpr:
		return notExpr(expr, env)
	case ast.NotEqualExpr:
		return notEqualExpr(expr, env)
	case ast.GtExpr:
		return gtExpr(expr, env)
	case ast.GteExpr:
		return gteExpr(expr, env)
	case ast.LtExpr:
		return ltExpr(expr, env)
	case ast.LteExpr:
		return lteExpr(expr, env)
	case ast.TernaryExpr:
		return ternaryExpr(expr, env)
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

func floatExpr(expr ast.FloatExpr) IType {
	i, err := strconv.ParseFloat(expr.Value, 64)
	if err != nil {
		panic(err)
	}
	return NewFloat(i)
}

func boolExpr(expr ast.BoolExpr) IType {
	return NewBool(expr.Value)
}

func nullExpr(expr ast.NullExpr) IType {
	return NewNull()
}

func ternaryExpr(expr ast.TernaryExpr, env *Env) IType {

	cond := expression(expr.Cond, env).Bool()
	if cond {
		return expression(expr.True, env)
	} else {
		return expression(expr.False, env)
	}
}

func mathExpr(expr ast.MathExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	switch expr.Operator {
	case "+":
		return Add(l, r)
	case "*":
		return Mul(l, r)
	case "-":
		return Sub(l, r)
	default:
		panic("not implemented" + expr.Operator)
	}

}

func identExpr(expr ast.IdentExpr, env *Env) IType {
	return env.Get(expr.Value)
}

func dotExpr(expr ast.DotExpr, env *Env) IType {
	// fmt.Println(expr)
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

func equalExpr(expr ast.EqualExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(l.Equal(r))
}

func gtExpr(expr ast.GtExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(l.Gt(r))
}

func gteExpr(expr ast.GteExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(l.Gte(r))
}
func ltExpr(expr ast.LtExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(l.Lt(r))
}
func lteExpr(expr ast.LteExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(l.Lte(r))
}

func notEqualExpr(expr ast.NotEqualExpr, env *Env) IType {
	l := expression(expr.Left, env)
	r := expression(expr.Right, env)
	return NewBool(!l.Equal(r))
}

func notExpr(expr ast.NotExpr, env *Env) IType {
	return NewBool(!expression(expr.Value, env).Bool())
}
