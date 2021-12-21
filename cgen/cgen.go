package cgen

import (
	"fmt"
	"strings"
)

type Context struct {
	result string
}

func (ctx *Context) AddHeader() {
	ctx.result += strings.Join([]string{
		"const print = (...args)=>console.log(...args)",
	}, "\n")
}

func CGen(stmts []interface{}) {
	ctx := Context{}
	ctx.AddHeader()

	fmt.Println(ExecStmts(stmts))

	// fmt.Println(ctx.result)
}
