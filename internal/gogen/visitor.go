package gogen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/threadedstream/tinycgo/internal/pkg/scope"
	parser "github.com/threadedstream/tinycgo/pkg/antlr"
	"strings"
)

type GogenVisitor struct {
	buf   strings.Builder
	ev    *EvalVisitor
	scope *scope.Scope
}

// NewGogenVisitor returns a fresh instance of GogenVisitor
func NewGogenVisitor() *GogenVisitor {
	return &GogenVisitor{
		buf: strings.Builder{},
	}
}

func (visitor *GogenVisitor) Visit(ctx antlr.Tree) any {
	switch child := ctx.(type) {
	default:
		panic("unknown ctx")
	case *parser.ProgramContext:
		return visitor.VisitProgram(child)
	case *parser.StatementContext:
		return visitor.VisitStatement(child)
	case *parser.IfNoElseStatementContext:
		return visitor.VisitIfNoElseStatement(child)
	case *parser.IfElseStatementContext:
		return visitor.VisitIfElseStatement(child)
	case *parser.BracedStatementContext:
		return visitor.VisitBracedStatement(child)
	case *parser.WhileStatementContext:
		return visitor.VisitWhileStatement(child)
	case *parser.DoWhileStatementContext:
		return visitor.VisitDoWhileStatement(child)
	case *parser.AssignmentExprContext:
		return visitor.VisitAssignmentExpr(child)
	case *parser.CompareSum_Context:
		return visitor.VisitCompareSum_(child)
	case *parser.Paren_exprContext:
		return visitor.VisitParen_expr(child)
	case *parser.ExprContext:
		return visitor.VisitExpr(child)
	case *parser.ExprSemiContext:
		return visitor.VisitExprSemi(child)
	case *parser.TestContext:
		return visitor.VisitTest(child)
	case *parser.Sum_Context:
		return visitor.VisitSum_(child)
	case *parser.TermContext:
		return visitor.VisitTerm(child)
	case *parser.Id_Context:
		return visitor.VisitId_(child)
	case *parser.IntegerContext:
		return visitor.VisitInteger(child)
	}
}

func (visitor *GogenVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	var results []any
	s := strings.Builder{}
	for _, s := range ctx.AllStatement() {
		results = append(results, visitor.Visit(s))
	}
	for _, r := range results {
		s.WriteString(r.(string))
	}
	return s.String()
}

func (visitor *GogenVisitor) VisitStatement(ctx *parser.StatementContext) any {
	if ctx.IfNoElseStatement() != nil {
		return visitor.Visit(ctx.IfNoElseStatement())
	} else if ctx.IfElseStatement() != nil {
		return visitor.Visit(ctx.IfElseStatement())
	} else if ctx.WhileStatement() != nil {
		return visitor.Visit(ctx.WhileStatement())
	} else if ctx.DoWhileStatement() != nil {
		return visitor.VisitDoWhileStatement(ctx.DoWhileStatement())
	} else if ctx.BracedStatement() != nil {
		return visitor.Visit(ctx.BracedStatement())
	} else if ctx.ExprSemi() != nil {
		return visitor.Visit(ctx.ExprSemi())
	} else if ctx.Semi() != nil {
		return ""
	}
	return nil
}

func (visitor *GogenVisitor) VisitIfNoElseStatement(stmt *parser.IfNoElseStatementContext) any {
	s := strings.Builder{}
	s.WriteString("if ")
	parenExprStr := visitor.Visit(stmt.Paren_expr()).(string)
	s.WriteString(parenExprStr + " != 0")
	stmtStr := visitor.Visit(stmt.Statement()).(string)
	s.WriteString(stmtStr)
	return s.String()
}

func (visitor *GogenVisitor) VisitIfElseStatement(stmt *parser.IfElseStatementContext) any {
	s := strings.Builder{}
	s.WriteString("if")
	s.WriteString(visitor.Visit(stmt.Paren_expr()).(string))
	s.WriteString(visitor.Visit(stmt.Statement(0)).(string))
	s.WriteString("else ")
	s.WriteString(visitor.Visit(stmt.Statement(1)).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitBracedStatement(stmt *parser.BracedStatementContext) any {
	scope.Push(&visitor.scope)
	s := strings.Builder{}
	s.WriteString("{ \n")
	for _, statement := range stmt.AllStatement() {
		s.WriteString(visitor.Visit(statement).(string))
	}
	s.WriteString("\n}\n")
	scope.Pop(&visitor.scope)
	return s.String()
}

func (visitor *GogenVisitor) VisitWhileStatement(stmt *parser.WhileStatementContext) any {
	s := strings.Builder{}
	s.WriteString("for ")
	s.WriteString(visitor.Visit(stmt.Paren_expr()).(string) + " {\n")
	s.WriteString(visitor.Visit(stmt.Statement()).(string))
	s.WriteString("\n}\n")
	return s.String()
}

func (visitor *GogenVisitor) VisitExprSemi(stmt *parser.ExprSemiContext) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(stmt.Expr()).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitDoWhileStatement(ctx parser.IDoWhileStatementContext) any {
	stmt := ctx.(*parser.DoWhileStatementContext)
	s := strings.Builder{}
	// at that point, one should rewrite tinyc's statement "do statement while paren_expr"
	// to golang's equivalent "for cond { }"
	// "cond" can easily be checked at transpile time, as expressions boil down to
	// binary operations applied to constants
	// if condition fails, we execute the body only once
	s.WriteString(visitor.Visit(stmt.Statement()).(string))
	s.WriteString("for ")
	s.WriteString(visitor.Visit(stmt.Paren_expr()).(string) + " ")
	s.WriteString(visitor.Visit(stmt.Statement()).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitParen_expr(ctx *parser.Paren_exprContext) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Expr()).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitExpr(ctx *parser.ExprContext) any {
	switch {
	case ctx.Test() != nil:
		return visitor.Visit(ctx.Test())
	case ctx.AssignmentExpr() != nil:
		return visitor.Visit(ctx.AssignmentExpr())
	}
	return nil
}

func (visitor *GogenVisitor) VisitTest(ctx *parser.TestContext) any {
	switch {
	case ctx.Sum_() != nil:
		return visitor.Visit(ctx.Sum_())
	case ctx.CompareSum_() != nil:
		return visitor.Visit(ctx.CompareSum_())
	}
	return nil
}

func (visitor *GogenVisitor) VisitAssignmentExpr(ctx *parser.AssignmentExprContext) any {
	s := strings.Builder{}
	s.WriteString(ctx.Id_().GetText())
	if _, ok := visitor.scope.Get(ctx.Id_().GetText()); ok {
		s.WriteString(" = ")
	} else {
		s.WriteString(" := ")
	}
	s.WriteString(visitor.Visit(ctx.Expr()).(string))
	s.WriteRune('\n')
	return s.String()
}

func (visitor *GogenVisitor) VisitCompareSum_(ctx *parser.CompareSum_Context) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Sum_(0)).(string))
	s.WriteString(" < ")
	s.WriteString(visitor.Visit(ctx.Sum_(1)).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitSum_(ctx *parser.Sum_Context) any {
	switch {
	default:
		return visitor.VisitBinarySum_(ctx)
	case ctx.Term() != nil && ctx.Sum_() == nil:
		return visitor.Visit(ctx.Term())
	}
}

func (visitor *GogenVisitor) VisitBinarySum_(ctx parser.ISum_Context) any {
	concreteCtx := ctx.(*parser.Sum_Context)
	binop := concreteCtx.Binop().(*parser.BinopContext)
	switch {
	case binop.PLUS() != nil:
		return visitor.VisitSum_Add(concreteCtx)
	case binop.MINUS() != nil:
		return visitor.VisitSum_Sub(concreteCtx)
	}
	return nil
}

func (visitor *GogenVisitor) VisitSum_Add(ctx *parser.Sum_Context) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Sum_()).(string))
	s.WriteString(" + ")
	s.WriteString(visitor.Visit(ctx.Term()).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitSum_Sub(ctx *parser.Sum_Context) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Sum_()).(string))
	s.WriteString(" - ")
	s.WriteString(visitor.Visit(ctx.Term()).(string))
	return s.String()
}

func (visitor *GogenVisitor) VisitTerm(ctx *parser.TermContext) any {
	switch {
	case ctx.Id_() != nil:
		return ctx.Id_().GetText()
	case ctx.Integer() != nil:
		return ctx.Integer().GetText()
	case ctx.Paren_expr() != nil:
		return visitor.Visit(ctx.Paren_expr())
	}
	return nil
}

func (visitor *GogenVisitor) VisitId_(ctx *parser.Id_Context) any {
	return ctx.STRING().GetText()
}

func (visitor *GogenVisitor) VisitInteger(ctx *parser.IntegerContext) any {
	return ctx.INT().GetText()
}

func (visitor *GogenVisitor) evaluateCondExpr(expr *parser.ExprContext) bool {
	return false
}
