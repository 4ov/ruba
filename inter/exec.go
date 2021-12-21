package inter

import (
	"io/ioutil"

	"github.com/4ov/ruba/ast"
	"github.com/4ov/ruba/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Options struct {
	Debug       bool
	AllowEval   bool
	AllowImport bool
}

type Ruba struct {
	Options Options
	Env     *Env
}

func (r *Ruba) DoString(raw string) error {
	_, err := r.execute(raw)
	return err
}

func (r *Ruba) DoFile(path string) error {
	bin_file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	_, err = r.execute(string(bin_file))
	return err
}

func (r *Ruba) execute(s string) (interface{}, error) {
	is := antlr.NewInputStream(s)
	lexer := parser.NewRubaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewRubaParser(stream)
	visitor := &ast.Visitor{}
	visitor.Program = ast.Program{Statements: []interface{}{}}
	program := p.Program().Accept(visitor).(ast.Program)

	return r.Exec(program.Statements, r.Env)
}

func (r *Ruba) Exec(stmts []interface{}, root_env *Env) (interface{}, error) {
	// var result interface{}
	// for _, stmt := range stmts {
	// 	switch stmt := stmt.(type) {
	// 	case ast.DeclareStmt:
	// 		result = declareStmt(stmt, root_env)
	// 	case ast.ExprStmt:
	// 		result = exprStmt(stmt, root_env)
	// 	default:
	// 		r.debug("unknown")
	// 	}
	// }

	return Exec(stmts, root_env), nil
}

func (r *Ruba) Eval(t string) (interface{}, error) {
	return r.execute(t)
}

func New() *Ruba {
	env := NewEnv()
	result := &Ruba{
		Env: env,
	}
	env.Ruba = result
	result.Options = Options{
		AllowEval:   true,
		AllowImport: true,
	}

	return result
}
