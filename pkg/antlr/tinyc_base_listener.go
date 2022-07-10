// Code generated from tinyc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tinyc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasetinycListener is a complete listener for a parse tree produced by tinycParser.
type BasetinycListener struct{}

var _ tinycListener = &BasetinycListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasetinycListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasetinycListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasetinycListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasetinycListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasetinycListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasetinycListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasetinycListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasetinycListener) ExitStatement(ctx *StatementContext) {}

// EnterIfNoElseStatement is called when production ifNoElseStatement is entered.
func (s *BasetinycListener) EnterIfNoElseStatement(ctx *IfNoElseStatementContext) {}

// ExitIfNoElseStatement is called when production ifNoElseStatement is exited.
func (s *BasetinycListener) ExitIfNoElseStatement(ctx *IfNoElseStatementContext) {}

// EnterIfElseStatement is called when production ifElseStatement is entered.
func (s *BasetinycListener) EnterIfElseStatement(ctx *IfElseStatementContext) {}

// ExitIfElseStatement is called when production ifElseStatement is exited.
func (s *BasetinycListener) ExitIfElseStatement(ctx *IfElseStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BasetinycListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BasetinycListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BasetinycListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BasetinycListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterBracedStatement is called when production bracedStatement is entered.
func (s *BasetinycListener) EnterBracedStatement(ctx *BracedStatementContext) {}

// ExitBracedStatement is called when production bracedStatement is exited.
func (s *BasetinycListener) ExitBracedStatement(ctx *BracedStatementContext) {}

// EnterExprSemi is called when production exprSemi is entered.
func (s *BasetinycListener) EnterExprSemi(ctx *ExprSemiContext) {}

// ExitExprSemi is called when production exprSemi is exited.
func (s *BasetinycListener) ExitExprSemi(ctx *ExprSemiContext) {}

// EnterSemi is called when production semi is entered.
func (s *BasetinycListener) EnterSemi(ctx *SemiContext) {}

// ExitSemi is called when production semi is exited.
func (s *BasetinycListener) ExitSemi(ctx *SemiContext) {}

// EnterParen_expr is called when production paren_expr is entered.
func (s *BasetinycListener) EnterParen_expr(ctx *Paren_exprContext) {}

// ExitParen_expr is called when production paren_expr is exited.
func (s *BasetinycListener) ExitParen_expr(ctx *Paren_exprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasetinycListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasetinycListener) ExitExpr(ctx *ExprContext) {}

// EnterAssignmentExpr is called when production assignmentExpr is entered.
func (s *BasetinycListener) EnterAssignmentExpr(ctx *AssignmentExprContext) {}

// ExitAssignmentExpr is called when production assignmentExpr is exited.
func (s *BasetinycListener) ExitAssignmentExpr(ctx *AssignmentExprContext) {}

// EnterTest is called when production test is entered.
func (s *BasetinycListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasetinycListener) ExitTest(ctx *TestContext) {}

// EnterCompareSum_ is called when production compareSum_ is entered.
func (s *BasetinycListener) EnterCompareSum_(ctx *CompareSum_Context) {}

// ExitCompareSum_ is called when production compareSum_ is exited.
func (s *BasetinycListener) ExitCompareSum_(ctx *CompareSum_Context) {}

// EnterSum_ is called when production sum_ is entered.
func (s *BasetinycListener) EnterSum_(ctx *Sum_Context) {}

// ExitSum_ is called when production sum_ is exited.
func (s *BasetinycListener) ExitSum_(ctx *Sum_Context) {}

// EnterBinop is called when production binop is entered.
func (s *BasetinycListener) EnterBinop(ctx *BinopContext) {}

// ExitBinop is called when production binop is exited.
func (s *BasetinycListener) ExitBinop(ctx *BinopContext) {}

// EnterTerm is called when production term is entered.
func (s *BasetinycListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BasetinycListener) ExitTerm(ctx *TermContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BasetinycListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BasetinycListener) ExitId_(ctx *Id_Context) {}

// EnterInteger is called when production integer is entered.
func (s *BasetinycListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasetinycListener) ExitInteger(ctx *IntegerContext) {}
