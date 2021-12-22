package main

import (
	"fmt"

	"github.com/4ov/ruba/inter"
)

func main() {

	r := inter.New()
	r.Options.Debug = true

	r.Env.Set("print", inter.NewNativeFn("print", func(env *inter.Env, args []inter.IType) inter.IType {
		var result []interface{}
		for _, arg := range args {
			result = append(result, arg.String())
		}
		fmt.Println(result...)
		return nil
	}))

	r.DoString(`
		
		
		
	`)

}

// func asPrintable(x []exec.Type) []interface{} {
// 	result := []interface{}{}
// 	for _, a := range x {
// 		switch a.TypeOf() {
// 		case exec.NUMBER, exec.BOOL, exec.NULL:
// 			result = append(result, color.YellowString(a.String()))
// 		case exec.FN:
// 			result = append(result, color.CyanString(a.String()))

// 		default:
// 			result = append(result, a)
// 		}
// 	}
// 	return result
// }

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fib(n-1) + fib(n-2)
// }

// type Int struct {
// 	Value int
// }
