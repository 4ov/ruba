package ast

import (
	"strings"

	"github.com/4ov/ruba/parser"
)

type Visitor struct {
	parser.BaseRubaVisitor
	Program Program
}

// func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
// 	return node.Accept(v)
// }

func (v *Visitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.Program.Statements = ctx.Block().Accept(v).([]interface{})
	return v.Program
}

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	result := []interface{}{}
	for _, stmt := range ctx.AllStmt() {
		result = append(result, stmt.Accept(v))
	}
	return result
}

//STMT
func (v *Visitor) VisitDeclareStmt(ctx *parser.DeclareStmtContext) interface{} {

	name := ctx.NAME.Accept(v)
	value := ctx.VALUE.Accept(v)
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()
	public := ctx.PUB != nil

	return DeclareStmt{
		Name:   name.(string),
		Value:  value,
		Line:   line,
		Public: public,
		Col:    col,
	}
}

func (v *Visitor) VisitImportStmt(ctx *parser.ImportStmtContext) interface{} {
	name := strings.Trim(ctx.NAME.Accept(v).(string), `"'`)
	var alias string
	if ctx.ALIAS != nil {
		alias = ctx.ALIAS.Accept(v).(string)
	}

	return ImportStmt{
		Name:  name,
		Alias: alias,
	}
}

func (v *Visitor) VisitAssignStmt(ctx *parser.AssignStmtContext) interface{} {
	name := ctx.ROOT.Accept(v)
	value := ctx.VALUE.Accept(v)
	chain := []Expression{}
	line := ctx.GetStart().GetLine()
	col := ctx.GetStart().GetColumn()
	for _, ch := range ctx.CHAIN {
		chain = append(chain, ch.Accept(v))
	}

	return AssignStmt{
		Name:  name.(string),
		Chain: chain,
		Value: value,
		Line:  line,
		Col:   col,
	}
}

func (v *Visitor) VisitFnCallStmt(ctx *parser.FnCallStmtContext) interface{} {
	return ctx.FnCall().Accept(v)
}

func (v *Visitor) VisitExprStmt(ctx *parser.ExprStmtContext) interface{} {
	return ExprStmt(ctx.E.Accept(v))
}

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	var result Expression
	if ctx.VALUE != nil {
		result = ctx.VALUE.Accept(v)
	}

	return ReturnStmt{
		Value: result,
	}
}

func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	var result Expression
	// if ctx.VALUE != nil {
	// 	result = ctx.VALUE.Accept(v)
	// }

	return BreakStmt{
		Value: result,
	}
}

func (v *Visitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	eifConds := []Expression{}
	eifBodies := [][]interface{}{}
	elseBody := []interface{}{}
	if ctx.ELIFCOND != nil {
		for _, cond := range ctx.ELIFCOND {
			eifConds = append(eifConds, cond.Accept(v))
		}
	}

	if ctx.ELIFBODY != nil {
		for _, body := range ctx.ELIFBODY {
			eifBodies = append(eifBodies, body.Accept(v).([]interface{}))
		}
	}

	if ctx.ELSEBODY != nil {
		elseBody = ctx.ELSEBODY.Accept(v).([]interface{})
	}

	return IfStmt{
		Cond:         ctx.COND.Accept(v),
		ElseIfConds:  eifConds,
		ElseIfBodies: eifBodies,
		ElseBody:     elseBody,
		Body:         ctx.BODY.Accept(v).([]interface{}),
	}
}

func (v *Visitor) VisitForInStmt(ctx *parser.ForInStmtContext) interface{} {
	names := []Expression{}
	for _, name := range ctx.NAMES {
		names = append(names, name.Accept(v))
	}
	target := ctx.TARGET.Accept(v)
	body := ctx.BODY.Accept(v)
	return ForInStmt{
		Names:  names,
		Target: target,
		Body:   body.([]interface{}),
	}
}

func (v *Visitor) VisitForCondStmt(ctx *parser.ForCondStmtContext) interface{} {

	cond := ctx.COND.Accept(v)
	body := ctx.BODY.Accept(v)
	var elseBody []interface{}
	if ctx.ELSEBODY != nil {
		elseBody = ctx.ELSEBODY.Accept(v).([]interface{})
	}
	return ForCondStmt{
		Cond:     cond,
		Body:     body.([]interface{}),
		ElseBody: elseBody,
	}
}

func (v *Visitor) VisitFnDeclareStmt(ctx *parser.FnDeclareStmtContext) interface{} {
	args := []Expression{}
	name := ctx.NAME.Accept(v)
	body := ctx.BODY.Accept(v)
	var restArg Expression
	for _, arg := range ctx.ARGS {
		args = append(args, arg.Accept(v))
	}

	if ctx.RESTARG != nil {
		restArg = ctx.RESTARG.Accept(v)
	}

	return FnDeclareStmt{
		Name:    name,
		Args:    args,
		RestArg: restArg,
		Body:    body.([]interface{}),
	}
}

//EXPR
func (v *Visitor) VisitIdentExpr(ctx *parser.IdentExprContext) interface{} {

	return IdentExpr{
		Value: ctx.Ident().Accept(v).(string),
	}
}

func (v *Visitor) VisitIncExpr(ctx *parser.IncExprContext) interface{} {

	return IncrExpr{
		Value: ctx.Expr().Accept(v),
	}
}

// func (v *Visitor) VisitIfExpr(ctx *parser.IfExprContext) interface{} {
// 	return IfExpr{
// 		Cond:  ctx.COND.Accept(v),
// 		True:  ctx.TRUE.Accept(v),
// 		False: ctx.FALSE.Accept(v),
// 	}
// }

func (v *Visitor) VisitGroupExpr(ctx *parser.GroupExprContext) interface{} {

	return ctx.Expr().Accept(v)
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) interface{} {
	return ctx.IntType().Accept(v)
}

func (v *Visitor) VisitFloatExpr(ctx *parser.FloatExprContext) interface{} {
	return ctx.FloatType().Accept(v)
}

func (v *Visitor) VisitNegativeExpr(ctx *parser.NegativeExprContext) interface{} {
	return NegativeExpr{
		Value: ctx.Expr().Accept(v),
	}
}

func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	result := strings.Trim(ctx.StringType().GetText(), `"'`)
	return StringExpr{
		Value: result,
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) interface{} {
	var result bool
	raw := ctx.GetText()
	if raw == "true" {
		result = true
	} else if raw == "false" {
		result = false
	}

	return BoolExpr{
		Value: result,
	}
}

func (v *Visitor) VisitNullExpr(ctx *parser.NullExprContext) interface{} {

	return NullExpr{}
}

func (v *Visitor) VisitArrayExpr(ctx *parser.ArrayExprContext) interface{} {
	elements := []Expression{}
	for _, element := range ctx.ELMS {
		elements = append(elements, element.Accept(v))
	}
	return ArrayExpr{
		Elements: elements,
	}
}

func (v *Visitor) VisitObjectExpr(ctx *parser.ObjectExprContext) interface{} {
	result := ObjectExpr{
		Elements: make(map[Expression]Expression),
	}
	for _, element := range ctx.ELMS {
		field := element.Accept(v).(ObjectField)
		result.Elements[field.Key] = field.Value
	}
	return result
}

func (v *Visitor) VisitMathExpr(ctx *parser.MathExprContext) interface{} {
	left := ctx.L.Accept(v).(Expression)
	right := ctx.R.Accept(v).(Expression)

	return MathExpr{
		Left:     left,
		Right:    right,
		Operator: ctx.MathSign().GetText(),
	}
}

func (v *Visitor) VisitEqualExpr(ctx *parser.EqualExprContext) interface{} {
	left := ctx.L.Accept(v).(Expression)
	right := ctx.R.Accept(v).(Expression)

	return EqualExpr{
		Left:  left,
		Right: right,
	}
}

func (v *Visitor) VisitNotEqualExpr(ctx *parser.NotEqualExprContext) interface{} {
	left := ctx.L.Accept(v).(Expression)
	right := ctx.R.Accept(v).(Expression)

	return NotEqualExpr{
		Left:  left,
		Right: right,
	}
}

func (v *Visitor) VisitGtExpr(ctx *parser.GtExprContext) interface{} {
	left := ctx.L.Accept(v).(Expression)
	right := ctx.R.Accept(v).(Expression)

	return GtExpr{
		Left:  left,
		Right: right,
	}
}

func (v *Visitor) VisitGteExpr(ctx *parser.GteExprContext) interface{} {
	left := ctx.L.Accept(v).(Expression)
	right := ctx.R.Accept(v).(Expression)

	return GteExpr{
		Left:  left,
		Right: right,
	}
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	return NotExpr{
		Value: ctx.Expr().Accept(v),
	}
}

func (v *Visitor) VisitDotExpr(ctx *parser.DotExprContext) interface{} {
	parent := ctx.PARENT.Accept(v).(Expression)
	child := ctx.CHILD.Accept(v).(Expression)

	return DotExpr{
		Parent: parent,
		Child:  child.(string),
	}
}

func (v *Visitor) VisitCallExpr(ctx *parser.CallExprContext) interface{} {
	target := ctx.NAME.Accept(v)
	args := []Expression{}
	for _, arg := range ctx.ARGS {
		args = append(args, arg.Accept(v))
	}

	return FnCallExpr{
		Target: target,
		Args:   args,
	}
}

func (v *Visitor) VisitIndexExpr(ctx *parser.IndexExprContext) interface{} {

	return IndexExpr{
		Parent: ctx.PARENT.Accept(v),
		Child:  ctx.CHILD.Accept(v),
	}
}

//types
func (v *Visitor) VisitIdent(ctx *parser.IdentContext) interface{} {

	return ctx.GetText()
}

func (v *Visitor) VisitStringType(ctx *parser.StringTypeContext) interface{} {
	return ctx.GetText()
}

func (v *Visitor) VisitIntType(ctx *parser.IntTypeContext) interface{} {
	return IntExpr{
		Value: ctx.GetText(),
	}
}

func (v *Visitor) VisitFloatType(ctx *parser.FloatTypeContext) interface{} {
	return FloatExpr{
		Value: ctx.GetText(),
	}
}

func (v *Visitor) VisitFnArg(ctx *parser.FnArgContext) interface{} {
	if ctx.ARG != nil {
		return ArgExpr{
			Value:      ctx.ARG.Accept(v),
			Spreadable: false,
		}
	} else {
		return ArgExpr{
			Value:      ctx.SPREADABLE.Accept(v),
			Spreadable: true,
		}
	}
}

func (v *Visitor) VisitObjectField(ctx *parser.ObjectFieldContext) interface{} {
	var name Expression
	value := ctx.VALUE.Accept(v)
	if ctx.STATIC != nil {
		name = StringExpr{
			Value: ctx.STATIC.GetText(),
		}
	} else {
		name = ctx.DYNAMIC.Accept(v)
	}
	return ObjectField{
		Key:   name,
		Value: value,
	}
}

func (v *Visitor) VisitFnCall(ctx *parser.FnCallContext) interface{} {
	var target Expression
	if ctx.NAME != nil {
		target = ctx.NAME.Accept(v)
	} else {
		target = ctx.EXPR.Accept(v)
	}

	args := []Expression{}
	for _, arg := range ctx.ARGS {
		args = append(args, arg.Accept(v))
	}
	return FnCallStmt{
		Target: target,
		Args:   args,
	}
}

func (v *Visitor) VisitNestable(ctx *parser.NestableContext) interface{} {
	if ctx.DOT != nil {

		return StringExpr{
			Value: ctx.DOT.Accept(v).(string),
		}
	} else { //INDEX
		return ctx.INDEX.Accept(v)
	}
}
