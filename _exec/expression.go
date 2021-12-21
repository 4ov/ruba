package exec

import (
	"fmt"
	"os"
	"plang/ast"
	"reflect"
	"strconv"
	"strings"
)

func expression_inner(expr ast.Expression, env *Env) Type {
	switch expr := expr.(type) {
	case ast.IntExpr:
		return intExpr(expr)
	case ast.FloatExpr:
		return floatExpr(expr)
	case ast.IdentExpr:
		return identExpr(expr, env)
	case ast.StringExpr:
		return stringExpr(expr, env)
	case ast.NegativeExpr:
		return negativeExpr(expr, env)
	case ast.BoolExpr:
		return boolExpr(expr)
	case ast.NullExpr:
		return nullExpr(expr)
	case ast.ArgExpr:
		return argExpr(expr, env)

	case ast.SubExpr:
		return subExpr(expr, env)
	case ast.DivExpr:
		return divExpr(expr, env)
	case ast.MulExpr:
		return mulExpr(expr, env)
	case ast.DotExpr:
		return dotExpr(expr, env)
	case ast.FnCallExpr:
		return callExpr(expr, env)
	case ast.ArrayExpr:
		return arrayExpr(expr, env)
	case ast.EqualExpr:
		return equalExpr(expr, env)
	case ast.NotEqualExpr:
		return notEqualExpr(expr, env)
	case ast.GtExpr:
		return gtExpr(expr, env)
	case ast.GteExpr:
		return gteExpr(expr, env)
	case ast.IncrExpr:
		return incrExpr(expr, env)
	case ast.MathExpr:
		return mathExpr(expr, env)
	default:
		fmt.Println("bad expression", expr, reflect.TypeOf(expr))
		os.Exit(1)
		return nil
	}
}

func expression(expr ast.Expression, env *Env) Type {

	result := expression_inner(expr, env)

	if value, ok := result.(*Getter); ok {
		return value.Value()
	} else {
		return result
	}

}

func intExpr(expr ast.IntExpr) Type {
	value, err := strconv.Atoi(expr.Value)
	if err != nil {
		fmt.Println("debug:", err)
	}
	return NewInt(value)
}

func floatExpr(expr ast.FloatExpr) Type {
	value, err := strconv.ParseFloat(expr.Value, 64)
	if err != nil {
		fmt.Println("debug:", err)
	}
	return NewFloat(value)
}

func identExpr(expr ast.IdentExpr, e *Env) Type {
	return e.GetVar(expr.Value)
}

func stringExpr(expr ast.StringExpr, e *Env) Type {
	return NewString(expr.Value)
}

func boolExpr(expr ast.BoolExpr) Type {
	return NewBool(expr.Value)
}

func nullExpr(expr ast.NullExpr) Type {
	return NewNull()
}

func negativeExpr(expr ast.NegativeExpr, e *Env) Type {
	value := expression(expr.Value, e)

	switch value := value.(type) {
	case *Int:
		return NewInt(-value.Value)
	case *Float:
		return NewFloat(-value.Value)
	default:
		NewError(fmt.Sprintf("Cannot use unray operator '-' on type %v", type_map[value.TypeOf()]))
		return nil
	}
}

//math
func mathExpr(expr ast.MathExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	fmt.Println(left, right)

	return NewInt(1)
}

func subExpr(expr ast.SubExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return left.(Number).Sub(right.(Number))
	default:
		NewError(fmt.Sprintf("Cannot substact values of type %v and %v", type_map[left.TypeOf()], type_map[right.TypeOf()]))

	}

	return NewNull()
}

func divExpr(expr ast.DivExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return left.(Number).Div(right.(Number))
	default:
		NewError(fmt.Sprintf("Cannot substact values of type %v and %v", type_map[left.TypeOf()], type_map[right.TypeOf()]))

	}

	return NewNull()
}

func mulExpr(expr ast.MulExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return left.(Number).Mul(right.(Number))
	default:
		NewError(fmt.Sprintf("Cannot substact values of type %v and %v", type_map[left.TypeOf()], type_map[right.TypeOf()]))

	}

	return NewNull()
}

//meta

func dotExpr(expr ast.DotExpr, env *Env) Type {
	root := expression(expr.Parent, env)
	return root.Access(expr.Child)
}

func incrExpr(expr ast.IncrExpr, env *Env) Type {

	switch target := expr.Value.(type) {
	case ast.IdentExpr:
		root := expression(target, env)
		env.SetVar(expr.Value.(ast.IdentExpr).Value, increaseValue(root))
		return root

	}

	return NullValue
}

func equalExpr(expr ast.EqualExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return left.(Number).Eq(right.(Number))
	case left.TypeOf() == STRING && right.TypeOf() == STRING:
		return NewBool(left.String() == right.String())
	case left.TypeOf() == NULL && right.TypeOf() == NULL:
		return TrueValue
	case left.TypeOf() == BOOL && right.TypeOf() == BOOL:
		return NewBool(left.(*Bool).Value == right.(*Bool).Value)
	default:
		return NewBool(false)
	}

}

func notEqualExpr(expr ast.NotEqualExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)

	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return NewBool(!left.(Number).Eq(right.(Number)).(*Bool).Value)
	case left.TypeOf() == STRING && right.TypeOf() == STRING:
		return NewBool(left.String() != right.String())
	case left.TypeOf() == NULL && right.TypeOf() == NULL:
		return NewBool(false)
	case left.TypeOf() == BOOL && right.TypeOf() == BOOL:
		return NewBool(left.(*Bool).Value != right.(*Bool).Value)
	default:
		return NewBool(true)
	}

}

func gtExpr(expr ast.GtExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)
	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return NewBool(left.(Number).Gt(right.(Number)).(*Bool).Value)
	case left.TypeOf() == STRING && right.TypeOf() == STRING:
		return NewBool(left.String() > right.String())
	case left.TypeOf() == NULL && right.TypeOf() == NULL:
		return NewBool(false)
	case left.TypeOf() == BOOL && right.TypeOf() == BOOL:
		return NewBool(false)
	default:
		return NewString("xx")
	}

}

func gteExpr(expr ast.GteExpr, env *Env) Type {
	left := expression(expr.Left, env)
	right := expression(expr.Right, env)
	switch {
	case left.TypeOf() == NUMBER && right.TypeOf() == NUMBER:
		return NewBool(left.(Number).Gte(right.(Number)).(*Bool).Value)
	case left.TypeOf() == STRING && right.TypeOf() == STRING:
		return NewBool(left.String() > right.String())
	case left.TypeOf() == NULL && right.TypeOf() == NULL:
		return NewBool(false)
	case left.TypeOf() == BOOL && right.TypeOf() == BOOL:
		return NewBool(false)
	default:
		return NewString("xx")
	}

}

func callExpr(expr ast.FnCallExpr, env *Env) Type {
	var args []Type
	for _, arg := range expr.Args {
		args = append(args, expression(arg, env))
	}
	target := expression(expr.Target, env)

	switch target := target.(type) {
	case *NativeFn:
		return CallNativeFn(*target, env, args)
	case *Fn:
		return CallFn(*target, env, args)
	default:
		NewError(fmt.Sprintf("Cannot call value of type %v", type_map[target.TypeOf()]))
	}

	return NewNull()
}

//arr && obj
func arrayExpr(expr ast.ArrayExpr, env *Env) Type {
	elms := []Type{}
	for _, elm := range expr.Elements {
		elms = append(elms, expression(elm, env))
	}

	return NewArray(elms)

}

func argExpr(expr ast.ArgExpr, env *Env) Type {
	if !expr.Spreadable {
		return expression(expr.Value, env)
	} else {
		result := expression(expr.Value, env)

		switch result := result.(type) {
		case *Array:
			return result
		case *String:
			r := []Type{}
			for _, c := range strings.Split(result.Value, "") {
				r = append(r, NewString(c))
			}
			return NewArray(r)
		default:
			NewError(fmt.Sprintf("Cannot spread value of type %v", type_map[result.TypeOf()]))
		}
	}

	return NewNull()

}
