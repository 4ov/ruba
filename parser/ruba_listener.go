// Generated from Ruba.g4 by ANTLR 4.7.

package parser // Ruba

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RubaListener is a complete listener for a parse tree produced by RubaParser.
type RubaListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// EnterIntType is called when entering the intType production.
	EnterIntType(c *IntTypeContext)

	// EnterFloatType is called when entering the floatType production.
	EnterFloatType(c *FloatTypeContext)

	// EnterStringType is called when entering the stringType production.
	EnterStringType(c *StringTypeContext)

	// EnterMathSign is called when entering the mathSign production.
	EnterMathSign(c *MathSignContext)

	// EnterNestable is called when entering the nestable production.
	EnterNestable(c *NestableContext)

	// EnterImportStmt is called when entering the ImportStmt production.
	EnterImportStmt(c *ImportStmtContext)

	// EnterDeclareStmt is called when entering the DeclareStmt production.
	EnterDeclareStmt(c *DeclareStmtContext)

	// EnterAssignStmt is called when entering the AssignStmt production.
	EnterAssignStmt(c *AssignStmtContext)

	// EnterFnCallStmt is called when entering the FnCallStmt production.
	EnterFnCallStmt(c *FnCallStmtContext)

	// EnterFnDeclareStmt is called when entering the FnDeclareStmt production.
	EnterFnDeclareStmt(c *FnDeclareStmtContext)

	// EnterForInStmt is called when entering the ForInStmt production.
	EnterForInStmt(c *ForInStmtContext)

	// EnterForCondStmt is called when entering the ForCondStmt production.
	EnterForCondStmt(c *ForCondStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterExprStmt is called when entering the ExprStmt production.
	EnterExprStmt(c *ExprStmtContext)

	// EnterFnArg is called when entering the fnArg production.
	EnterFnArg(c *FnArgContext)

	// EnterFnCall is called when entering the fnCall production.
	EnterFnCall(c *FnCallContext)

	// EnterObjectField is called when entering the objectField production.
	EnterObjectField(c *ObjectFieldContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterStringExpr is called when entering the StringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterEqualExpr is called when entering the EqualExpr production.
	EnterEqualExpr(c *EqualExprContext)

	// EnterIdentExpr is called when entering the IdentExpr production.
	EnterIdentExpr(c *IdentExprContext)

	// EnterDecExpr is called when entering the DecExpr production.
	EnterDecExpr(c *DecExprContext)

	// EnterFloatExpr is called when entering the FloatExpr production.
	EnterFloatExpr(c *FloatExprContext)

	// EnterIncExpr is called when entering the IncExpr production.
	EnterIncExpr(c *IncExprContext)

	// EnterAddOrSubExpr is called when entering the AddOrSubExpr production.
	EnterAddOrSubExpr(c *AddOrSubExprContext)

	// EnterGtExpr is called when entering the GtExpr production.
	EnterGtExpr(c *GtExprContext)

	// EnterLtExpr is called when entering the LtExpr production.
	EnterLtExpr(c *LtExprContext)

	// EnterDotExpr is called when entering the DotExpr production.
	EnterDotExpr(c *DotExprContext)

	// EnterNotEqualExpr is called when entering the NotEqualExpr production.
	EnterNotEqualExpr(c *NotEqualExprContext)

	// EnterLteExpr is called when entering the LteExpr production.
	EnterLteExpr(c *LteExprContext)

	// EnterIndexExpr is called when entering the IndexExpr production.
	EnterIndexExpr(c *IndexExprContext)

	// EnterNullExpr is called when entering the NullExpr production.
	EnterNullExpr(c *NullExprContext)

	// EnterGteExpr is called when entering the GteExpr production.
	EnterGteExpr(c *GteExprContext)

	// EnterGroupExpr is called when entering the GroupExpr production.
	EnterGroupExpr(c *GroupExprContext)

	// EnterNegativeExpr is called when entering the NegativeExpr production.
	EnterNegativeExpr(c *NegativeExprContext)

	// EnterArrayExpr is called when entering the ArrayExpr production.
	EnterArrayExpr(c *ArrayExprContext)

	// EnterObjectExpr is called when entering the ObjectExpr production.
	EnterObjectExpr(c *ObjectExprContext)

	// EnterMulOrDivExpr is called when entering the MulOrDivExpr production.
	EnterMulOrDivExpr(c *MulOrDivExprContext)

	// EnterCallExpr is called when entering the CallExpr production.
	EnterCallExpr(c *CallExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)

	// ExitIntType is called when exiting the intType production.
	ExitIntType(c *IntTypeContext)

	// ExitFloatType is called when exiting the floatType production.
	ExitFloatType(c *FloatTypeContext)

	// ExitStringType is called when exiting the stringType production.
	ExitStringType(c *StringTypeContext)

	// ExitMathSign is called when exiting the mathSign production.
	ExitMathSign(c *MathSignContext)

	// ExitNestable is called when exiting the nestable production.
	ExitNestable(c *NestableContext)

	// ExitImportStmt is called when exiting the ImportStmt production.
	ExitImportStmt(c *ImportStmtContext)

	// ExitDeclareStmt is called when exiting the DeclareStmt production.
	ExitDeclareStmt(c *DeclareStmtContext)

	// ExitAssignStmt is called when exiting the AssignStmt production.
	ExitAssignStmt(c *AssignStmtContext)

	// ExitFnCallStmt is called when exiting the FnCallStmt production.
	ExitFnCallStmt(c *FnCallStmtContext)

	// ExitFnDeclareStmt is called when exiting the FnDeclareStmt production.
	ExitFnDeclareStmt(c *FnDeclareStmtContext)

	// ExitForInStmt is called when exiting the ForInStmt production.
	ExitForInStmt(c *ForInStmtContext)

	// ExitForCondStmt is called when exiting the ForCondStmt production.
	ExitForCondStmt(c *ForCondStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitExprStmt is called when exiting the ExprStmt production.
	ExitExprStmt(c *ExprStmtContext)

	// ExitFnArg is called when exiting the fnArg production.
	ExitFnArg(c *FnArgContext)

	// ExitFnCall is called when exiting the fnCall production.
	ExitFnCall(c *FnCallContext)

	// ExitObjectField is called when exiting the objectField production.
	ExitObjectField(c *ObjectFieldContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitStringExpr is called when exiting the StringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitEqualExpr is called when exiting the EqualExpr production.
	ExitEqualExpr(c *EqualExprContext)

	// ExitIdentExpr is called when exiting the IdentExpr production.
	ExitIdentExpr(c *IdentExprContext)

	// ExitDecExpr is called when exiting the DecExpr production.
	ExitDecExpr(c *DecExprContext)

	// ExitFloatExpr is called when exiting the FloatExpr production.
	ExitFloatExpr(c *FloatExprContext)

	// ExitIncExpr is called when exiting the IncExpr production.
	ExitIncExpr(c *IncExprContext)

	// ExitAddOrSubExpr is called when exiting the AddOrSubExpr production.
	ExitAddOrSubExpr(c *AddOrSubExprContext)

	// ExitGtExpr is called when exiting the GtExpr production.
	ExitGtExpr(c *GtExprContext)

	// ExitLtExpr is called when exiting the LtExpr production.
	ExitLtExpr(c *LtExprContext)

	// ExitDotExpr is called when exiting the DotExpr production.
	ExitDotExpr(c *DotExprContext)

	// ExitNotEqualExpr is called when exiting the NotEqualExpr production.
	ExitNotEqualExpr(c *NotEqualExprContext)

	// ExitLteExpr is called when exiting the LteExpr production.
	ExitLteExpr(c *LteExprContext)

	// ExitIndexExpr is called when exiting the IndexExpr production.
	ExitIndexExpr(c *IndexExprContext)

	// ExitNullExpr is called when exiting the NullExpr production.
	ExitNullExpr(c *NullExprContext)

	// ExitGteExpr is called when exiting the GteExpr production.
	ExitGteExpr(c *GteExprContext)

	// ExitGroupExpr is called when exiting the GroupExpr production.
	ExitGroupExpr(c *GroupExprContext)

	// ExitNegativeExpr is called when exiting the NegativeExpr production.
	ExitNegativeExpr(c *NegativeExprContext)

	// ExitArrayExpr is called when exiting the ArrayExpr production.
	ExitArrayExpr(c *ArrayExprContext)

	// ExitObjectExpr is called when exiting the ObjectExpr production.
	ExitObjectExpr(c *ObjectExprContext)

	// ExitMulOrDivExpr is called when exiting the MulOrDivExpr production.
	ExitMulOrDivExpr(c *MulOrDivExprContext)

	// ExitCallExpr is called when exiting the CallExpr production.
	ExitCallExpr(c *CallExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)
}
