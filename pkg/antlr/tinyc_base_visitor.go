// Code generated from tinyc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // tinyc

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasetinycVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetinycVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitIfNoElseStatement(ctx *IfNoElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitIfElseStatement(ctx *IfElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitBracedStatement(ctx *BracedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitExprSemi(ctx *ExprSemiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitSemi(ctx *SemiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitParen_expr(ctx *Paren_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitAssignmentExpr(ctx *AssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitCompareSum_(ctx *CompareSum_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitSum_(ctx *Sum_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitBinop(ctx *BinopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitId_(ctx *Id_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetinycVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}
