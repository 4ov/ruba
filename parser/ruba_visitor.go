// Generated from Ruba.g4 by ANTLR 4.7.

package parser // Ruba

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by RubaParser.
type RubaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RubaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RubaParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by RubaParser#ident.
	VisitIdent(ctx *IdentContext) interface{}

	// Visit a parse tree produced by RubaParser#intType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by RubaParser#floatType.
	VisitFloatType(ctx *FloatTypeContext) interface{}

	// Visit a parse tree produced by RubaParser#stringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by RubaParser#mathSign.
	VisitMathSign(ctx *MathSignContext) interface{}

	// Visit a parse tree produced by RubaParser#nestable.
	VisitNestable(ctx *NestableContext) interface{}

	// Visit a parse tree produced by RubaParser#ImportStmt.
	VisitImportStmt(ctx *ImportStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#DeclareStmt.
	VisitDeclareStmt(ctx *DeclareStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#AssignStmt.
	VisitAssignStmt(ctx *AssignStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#FnCallStmt.
	VisitFnCallStmt(ctx *FnCallStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#FnDeclareStmt.
	VisitFnDeclareStmt(ctx *FnDeclareStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#ForInStmt.
	VisitForInStmt(ctx *ForInStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#ForCondStmt.
	VisitForCondStmt(ctx *ForCondStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#ExprStmt.
	VisitExprStmt(ctx *ExprStmtContext) interface{}

	// Visit a parse tree produced by RubaParser#fnArg.
	VisitFnArg(ctx *FnArgContext) interface{}

	// Visit a parse tree produced by RubaParser#fnCall.
	VisitFnCall(ctx *FnCallContext) interface{}

	// Visit a parse tree produced by RubaParser#objectField.
	VisitObjectField(ctx *ObjectFieldContext) interface{}

	// Visit a parse tree produced by RubaParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by RubaParser#StringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by RubaParser#EqualExpr.
	VisitEqualExpr(ctx *EqualExprContext) interface{}

	// Visit a parse tree produced by RubaParser#IdentExpr.
	VisitIdentExpr(ctx *IdentExprContext) interface{}

	// Visit a parse tree produced by RubaParser#DecExpr.
	VisitDecExpr(ctx *DecExprContext) interface{}

	// Visit a parse tree produced by RubaParser#FloatExpr.
	VisitFloatExpr(ctx *FloatExprContext) interface{}

	// Visit a parse tree produced by RubaParser#IncExpr.
	VisitIncExpr(ctx *IncExprContext) interface{}

	// Visit a parse tree produced by RubaParser#GtExpr.
	VisitGtExpr(ctx *GtExprContext) interface{}

	// Visit a parse tree produced by RubaParser#LtExpr.
	VisitLtExpr(ctx *LtExprContext) interface{}

	// Visit a parse tree produced by RubaParser#DotExpr.
	VisitDotExpr(ctx *DotExprContext) interface{}

	// Visit a parse tree produced by RubaParser#NotEqualExpr.
	VisitNotEqualExpr(ctx *NotEqualExprContext) interface{}

	// Visit a parse tree produced by RubaParser#LteExpr.
	VisitLteExpr(ctx *LteExprContext) interface{}

	// Visit a parse tree produced by RubaParser#IndexExpr.
	VisitIndexExpr(ctx *IndexExprContext) interface{}

	// Visit a parse tree produced by RubaParser#NullExpr.
	VisitNullExpr(ctx *NullExprContext) interface{}

	// Visit a parse tree produced by RubaParser#GteExpr.
	VisitGteExpr(ctx *GteExprContext) interface{}

	// Visit a parse tree produced by RubaParser#GroupExpr.
	VisitGroupExpr(ctx *GroupExprContext) interface{}

	// Visit a parse tree produced by RubaParser#NegativeExpr.
	VisitNegativeExpr(ctx *NegativeExprContext) interface{}

	// Visit a parse tree produced by RubaParser#ArrayExpr.
	VisitArrayExpr(ctx *ArrayExprContext) interface{}

	// Visit a parse tree produced by RubaParser#ObjectExpr.
	VisitObjectExpr(ctx *ObjectExprContext) interface{}

	// Visit a parse tree produced by RubaParser#CallExpr.
	VisitCallExpr(ctx *CallExprContext) interface{}

	// Visit a parse tree produced by RubaParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by RubaParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by RubaParser#MathExpr.
	VisitMathExpr(ctx *MathExprContext) interface{}
}
