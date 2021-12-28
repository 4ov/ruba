// Generated from Ruba.g4 by ANTLR 4.7.

package parser // Ruba

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRubaListener is a complete listener for a parse tree produced by RubaParser.
type BaseRubaListener struct{}

var _ RubaListener = &BaseRubaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRubaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRubaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRubaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRubaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRubaListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRubaListener) ExitProgram(ctx *ProgramContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseRubaListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseRubaListener) ExitBlock(ctx *BlockContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseRubaListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseRubaListener) ExitIdent(ctx *IdentContext) {}

// EnterIntType is called when production intType is entered.
func (s *BaseRubaListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production intType is exited.
func (s *BaseRubaListener) ExitIntType(ctx *IntTypeContext) {}

// EnterFloatType is called when production floatType is entered.
func (s *BaseRubaListener) EnterFloatType(ctx *FloatTypeContext) {}

// ExitFloatType is called when production floatType is exited.
func (s *BaseRubaListener) ExitFloatType(ctx *FloatTypeContext) {}

// EnterStringType is called when production stringType is entered.
func (s *BaseRubaListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production stringType is exited.
func (s *BaseRubaListener) ExitStringType(ctx *StringTypeContext) {}

// EnterMathSign is called when production mathSign is entered.
func (s *BaseRubaListener) EnterMathSign(ctx *MathSignContext) {}

// ExitMathSign is called when production mathSign is exited.
func (s *BaseRubaListener) ExitMathSign(ctx *MathSignContext) {}

// EnterNestable is called when production nestable is entered.
func (s *BaseRubaListener) EnterNestable(ctx *NestableContext) {}

// ExitNestable is called when production nestable is exited.
func (s *BaseRubaListener) ExitNestable(ctx *NestableContext) {}

// EnterImportStmt is called when production ImportStmt is entered.
func (s *BaseRubaListener) EnterImportStmt(ctx *ImportStmtContext) {}

// ExitImportStmt is called when production ImportStmt is exited.
func (s *BaseRubaListener) ExitImportStmt(ctx *ImportStmtContext) {}

// EnterDeclareStmt is called when production DeclareStmt is entered.
func (s *BaseRubaListener) EnterDeclareStmt(ctx *DeclareStmtContext) {}

// ExitDeclareStmt is called when production DeclareStmt is exited.
func (s *BaseRubaListener) ExitDeclareStmt(ctx *DeclareStmtContext) {}

// EnterAssignStmt is called when production AssignStmt is entered.
func (s *BaseRubaListener) EnterAssignStmt(ctx *AssignStmtContext) {}

// ExitAssignStmt is called when production AssignStmt is exited.
func (s *BaseRubaListener) ExitAssignStmt(ctx *AssignStmtContext) {}

// EnterFnCallStmt is called when production FnCallStmt is entered.
func (s *BaseRubaListener) EnterFnCallStmt(ctx *FnCallStmtContext) {}

// ExitFnCallStmt is called when production FnCallStmt is exited.
func (s *BaseRubaListener) ExitFnCallStmt(ctx *FnCallStmtContext) {}

// EnterFnDeclareStmt is called when production FnDeclareStmt is entered.
func (s *BaseRubaListener) EnterFnDeclareStmt(ctx *FnDeclareStmtContext) {}

// ExitFnDeclareStmt is called when production FnDeclareStmt is exited.
func (s *BaseRubaListener) ExitFnDeclareStmt(ctx *FnDeclareStmtContext) {}

// EnterForInStmt is called when production ForInStmt is entered.
func (s *BaseRubaListener) EnterForInStmt(ctx *ForInStmtContext) {}

// ExitForInStmt is called when production ForInStmt is exited.
func (s *BaseRubaListener) ExitForInStmt(ctx *ForInStmtContext) {}

// EnterForCondStmt is called when production ForCondStmt is entered.
func (s *BaseRubaListener) EnterForCondStmt(ctx *ForCondStmtContext) {}

// ExitForCondStmt is called when production ForCondStmt is exited.
func (s *BaseRubaListener) ExitForCondStmt(ctx *ForCondStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseRubaListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseRubaListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseRubaListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseRubaListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseRubaListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseRubaListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterExprStmt is called when production ExprStmt is entered.
func (s *BaseRubaListener) EnterExprStmt(ctx *ExprStmtContext) {}

// ExitExprStmt is called when production ExprStmt is exited.
func (s *BaseRubaListener) ExitExprStmt(ctx *ExprStmtContext) {}

// EnterFnArg is called when production fnArg is entered.
func (s *BaseRubaListener) EnterFnArg(ctx *FnArgContext) {}

// ExitFnArg is called when production fnArg is exited.
func (s *BaseRubaListener) ExitFnArg(ctx *FnArgContext) {}

// EnterFnCall is called when production fnCall is entered.
func (s *BaseRubaListener) EnterFnCall(ctx *FnCallContext) {}

// ExitFnCall is called when production fnCall is exited.
func (s *BaseRubaListener) ExitFnCall(ctx *FnCallContext) {}

// EnterObjectField is called when production objectField is entered.
func (s *BaseRubaListener) EnterObjectField(ctx *ObjectFieldContext) {}

// ExitObjectField is called when production objectField is exited.
func (s *BaseRubaListener) ExitObjectField(ctx *ObjectFieldContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseRubaListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseRubaListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterStringExpr is called when production StringExpr is entered.
func (s *BaseRubaListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production StringExpr is exited.
func (s *BaseRubaListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterEqualExpr is called when production EqualExpr is entered.
func (s *BaseRubaListener) EnterEqualExpr(ctx *EqualExprContext) {}

// ExitEqualExpr is called when production EqualExpr is exited.
func (s *BaseRubaListener) ExitEqualExpr(ctx *EqualExprContext) {}

// EnterFloatExpr is called when production FloatExpr is entered.
func (s *BaseRubaListener) EnterFloatExpr(ctx *FloatExprContext) {}

// ExitFloatExpr is called when production FloatExpr is exited.
func (s *BaseRubaListener) ExitFloatExpr(ctx *FloatExprContext) {}

// EnterGtExpr is called when production GtExpr is entered.
func (s *BaseRubaListener) EnterGtExpr(ctx *GtExprContext) {}

// ExitGtExpr is called when production GtExpr is exited.
func (s *BaseRubaListener) ExitGtExpr(ctx *GtExprContext) {}

// EnterLtExpr is called when production LtExpr is entered.
func (s *BaseRubaListener) EnterLtExpr(ctx *LtExprContext) {}

// ExitLtExpr is called when production LtExpr is exited.
func (s *BaseRubaListener) ExitLtExpr(ctx *LtExprContext) {}

// EnterIndexExpr is called when production IndexExpr is entered.
func (s *BaseRubaListener) EnterIndexExpr(ctx *IndexExprContext) {}

// ExitIndexExpr is called when production IndexExpr is exited.
func (s *BaseRubaListener) ExitIndexExpr(ctx *IndexExprContext) {}

// EnterNullExpr is called when production NullExpr is entered.
func (s *BaseRubaListener) EnterNullExpr(ctx *NullExprContext) {}

// ExitNullExpr is called when production NullExpr is exited.
func (s *BaseRubaListener) ExitNullExpr(ctx *NullExprContext) {}

// EnterObjectExpr is called when production ObjectExpr is entered.
func (s *BaseRubaListener) EnterObjectExpr(ctx *ObjectExprContext) {}

// ExitObjectExpr is called when production ObjectExpr is exited.
func (s *BaseRubaListener) ExitObjectExpr(ctx *ObjectExprContext) {}

// EnterCallExpr is called when production CallExpr is entered.
func (s *BaseRubaListener) EnterCallExpr(ctx *CallExprContext) {}

// ExitCallExpr is called when production CallExpr is exited.
func (s *BaseRubaListener) ExitCallExpr(ctx *CallExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseRubaListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseRubaListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterTernaryExpr is called when production TernaryExpr is entered.
func (s *BaseRubaListener) EnterTernaryExpr(ctx *TernaryExprContext) {}

// ExitTernaryExpr is called when production TernaryExpr is exited.
func (s *BaseRubaListener) ExitTernaryExpr(ctx *TernaryExprContext) {}

// EnterIdentExpr is called when production IdentExpr is entered.
func (s *BaseRubaListener) EnterIdentExpr(ctx *IdentExprContext) {}

// ExitIdentExpr is called when production IdentExpr is exited.
func (s *BaseRubaListener) ExitIdentExpr(ctx *IdentExprContext) {}

// EnterDecExpr is called when production DecExpr is entered.
func (s *BaseRubaListener) EnterDecExpr(ctx *DecExprContext) {}

// ExitDecExpr is called when production DecExpr is exited.
func (s *BaseRubaListener) ExitDecExpr(ctx *DecExprContext) {}

// EnterIncExpr is called when production IncExpr is entered.
func (s *BaseRubaListener) EnterIncExpr(ctx *IncExprContext) {}

// ExitIncExpr is called when production IncExpr is exited.
func (s *BaseRubaListener) ExitIncExpr(ctx *IncExprContext) {}

// EnterAddOrSubExpr is called when production AddOrSubExpr is entered.
func (s *BaseRubaListener) EnterAddOrSubExpr(ctx *AddOrSubExprContext) {}

// ExitAddOrSubExpr is called when production AddOrSubExpr is exited.
func (s *BaseRubaListener) ExitAddOrSubExpr(ctx *AddOrSubExprContext) {}

// EnterDotExpr is called when production DotExpr is entered.
func (s *BaseRubaListener) EnterDotExpr(ctx *DotExprContext) {}

// ExitDotExpr is called when production DotExpr is exited.
func (s *BaseRubaListener) ExitDotExpr(ctx *DotExprContext) {}

// EnterNotEqualExpr is called when production NotEqualExpr is entered.
func (s *BaseRubaListener) EnterNotEqualExpr(ctx *NotEqualExprContext) {}

// ExitNotEqualExpr is called when production NotEqualExpr is exited.
func (s *BaseRubaListener) ExitNotEqualExpr(ctx *NotEqualExprContext) {}

// EnterLteExpr is called when production LteExpr is entered.
func (s *BaseRubaListener) EnterLteExpr(ctx *LteExprContext) {}

// ExitLteExpr is called when production LteExpr is exited.
func (s *BaseRubaListener) ExitLteExpr(ctx *LteExprContext) {}

// EnterGteExpr is called when production GteExpr is entered.
func (s *BaseRubaListener) EnterGteExpr(ctx *GteExprContext) {}

// ExitGteExpr is called when production GteExpr is exited.
func (s *BaseRubaListener) ExitGteExpr(ctx *GteExprContext) {}

// EnterGroupExpr is called when production GroupExpr is entered.
func (s *BaseRubaListener) EnterGroupExpr(ctx *GroupExprContext) {}

// ExitGroupExpr is called when production GroupExpr is exited.
func (s *BaseRubaListener) ExitGroupExpr(ctx *GroupExprContext) {}

// EnterNegativeExpr is called when production NegativeExpr is entered.
func (s *BaseRubaListener) EnterNegativeExpr(ctx *NegativeExprContext) {}

// ExitNegativeExpr is called when production NegativeExpr is exited.
func (s *BaseRubaListener) ExitNegativeExpr(ctx *NegativeExprContext) {}

// EnterArrayExpr is called when production ArrayExpr is entered.
func (s *BaseRubaListener) EnterArrayExpr(ctx *ArrayExprContext) {}

// ExitArrayExpr is called when production ArrayExpr is exited.
func (s *BaseRubaListener) ExitArrayExpr(ctx *ArrayExprContext) {}

// EnterMulOrDivExpr is called when production MulOrDivExpr is entered.
func (s *BaseRubaListener) EnterMulOrDivExpr(ctx *MulOrDivExprContext) {}

// ExitMulOrDivExpr is called when production MulOrDivExpr is exited.
func (s *BaseRubaListener) ExitMulOrDivExpr(ctx *MulOrDivExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseRubaListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseRubaListener) ExitIntExpr(ctx *IntExprContext) {}
