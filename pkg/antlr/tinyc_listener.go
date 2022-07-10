// Code generated from tinyc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tinyc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// tinycListener is a complete listener for a parse tree produced by tinycParser.
type tinycListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterIfNoElseStatement is called when entering the ifNoElseStatement production.
	EnterIfNoElseStatement(c *IfNoElseStatementContext)

	// EnterIfElseStatement is called when entering the ifElseStatement production.
	EnterIfElseStatement(c *IfElseStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterBracedStatement is called when entering the bracedStatement production.
	EnterBracedStatement(c *BracedStatementContext)

	// EnterExprSemi is called when entering the exprSemi production.
	EnterExprSemi(c *ExprSemiContext)

	// EnterSemi is called when entering the semi production.
	EnterSemi(c *SemiContext)

	// EnterParen_expr is called when entering the paren_expr production.
	EnterParen_expr(c *Paren_exprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAssignmentExpr is called when entering the assignmentExpr production.
	EnterAssignmentExpr(c *AssignmentExprContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterCompareSum_ is called when entering the compareSum_ production.
	EnterCompareSum_(c *CompareSum_Context)

	// EnterSum_ is called when entering the sum_ production.
	EnterSum_(c *Sum_Context)

	// EnterBinop is called when entering the binop production.
	EnterBinop(c *BinopContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitIfNoElseStatement is called when exiting the ifNoElseStatement production.
	ExitIfNoElseStatement(c *IfNoElseStatementContext)

	// ExitIfElseStatement is called when exiting the ifElseStatement production.
	ExitIfElseStatement(c *IfElseStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitBracedStatement is called when exiting the bracedStatement production.
	ExitBracedStatement(c *BracedStatementContext)

	// ExitExprSemi is called when exiting the exprSemi production.
	ExitExprSemi(c *ExprSemiContext)

	// ExitSemi is called when exiting the semi production.
	ExitSemi(c *SemiContext)

	// ExitParen_expr is called when exiting the paren_expr production.
	ExitParen_expr(c *Paren_exprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAssignmentExpr is called when exiting the assignmentExpr production.
	ExitAssignmentExpr(c *AssignmentExprContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitCompareSum_ is called when exiting the compareSum_ production.
	ExitCompareSum_(c *CompareSum_Context)

	// ExitSum_ is called when exiting the sum_ production.
	ExitSum_(c *Sum_Context)

	// ExitBinop is called when exiting the binop production.
	ExitBinop(c *BinopContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)
}
