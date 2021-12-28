// Generated from Ruba.g4 by ANTLR 4.7.

package parser // Ruba

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseRubaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRubaVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFloatType(ctx *FloatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitMathSign(ctx *MathSignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitNestable(ctx *NestableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitImportStmt(ctx *ImportStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitDeclareStmt(ctx *DeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitAssignStmt(ctx *AssignStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFnCallStmt(ctx *FnCallStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFnDeclareStmt(ctx *FnDeclareStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitForInStmt(ctx *ForInStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitForCondStmt(ctx *ForCondStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitExprStmt(ctx *ExprStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFnArg(ctx *FnArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFnCall(ctx *FnCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitObjectField(ctx *ObjectFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitEqualExpr(ctx *EqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitGtExpr(ctx *GtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitLtExpr(ctx *LtExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIndexExpr(ctx *IndexExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitNullExpr(ctx *NullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitObjectExpr(ctx *ObjectExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitCallExpr(ctx *CallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitTernaryExpr(ctx *TernaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIdentExpr(ctx *IdentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitDecExpr(ctx *DecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIncExpr(ctx *IncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitAddOrSubExpr(ctx *AddOrSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitDotExpr(ctx *DotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitNotEqualExpr(ctx *NotEqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitLteExpr(ctx *LteExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitGteExpr(ctx *GteExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitGroupExpr(ctx *GroupExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitNegativeExpr(ctx *NegativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitArrayExpr(ctx *ArrayExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitMulOrDivExpr(ctx *MulOrDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRubaVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}
