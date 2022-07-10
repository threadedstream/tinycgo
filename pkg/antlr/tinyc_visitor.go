// Code generated from tinyc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tinyc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by tinycParser.
type tinycVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tinycParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tinycParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tinycParser#ifNoElseStatement.
	VisitIfNoElseStatement(ctx *IfNoElseStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#ifElseStatement.
	VisitIfElseStatement(ctx *IfElseStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#doWhileStatement.
	VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#bracedStatement.
	VisitBracedStatement(ctx *BracedStatementContext) interface{}

	// Visit a parse tree produced by tinycParser#exprSemi.
	VisitExprSemi(ctx *ExprSemiContext) interface{}

	// Visit a parse tree produced by tinycParser#semi.
	VisitSemi(ctx *SemiContext) interface{}

	// Visit a parse tree produced by tinycParser#paren_expr.
	VisitParen_expr(ctx *Paren_exprContext) interface{}

	// Visit a parse tree produced by tinycParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by tinycParser#assignmentExpr.
	VisitAssignmentExpr(ctx *AssignmentExprContext) interface{}

	// Visit a parse tree produced by tinycParser#test.
	VisitTest(ctx *TestContext) interface{}

	// Visit a parse tree produced by tinycParser#compareSum_.
	VisitCompareSum_(ctx *CompareSum_Context) interface{}

	// Visit a parse tree produced by tinycParser#sum_.
	VisitSum_(ctx *Sum_Context) interface{}

	// Visit a parse tree produced by tinycParser#binop.
	VisitBinop(ctx *BinopContext) interface{}

	// Visit a parse tree produced by tinycParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by tinycParser#id_.
	VisitId_(ctx *Id_Context) interface{}

	// Visit a parse tree produced by tinycParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}
}
