package jsgen

import (
	"fmt"
	"plang/ast"
	"strings"
)

func constantHeader() string {
	lines := []string{
		"let print = (...args) => { console.log(...args) };",
		"let panic = (message) => { throw message };",
		"const CONSTS = { platform : 'js' }",
	}

	return fmt.Sprintln(strings.Join(lines, "\n"))
}

func Generate(program ast.Program) string {
	result := ``
	result += constantHeader()
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case ast.DeclareStmt:
			result += declareStmt(stmt)
		case ast.FnCallStmt:
			result += fnCallStmt(stmt)
		}
	}
	return result
}
