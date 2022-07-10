package gogen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/threadedstream/tinycgo/pkg/antlr"
	"strings"
	"fmt"
)

type GogenVisitor struct {
	buf strings.Builder
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
	case *parser.Paren_exprContext:
		return visitor.VisitParen_expr(child)
	case *parser.ExprContext:
		return visitor.VisitExpr(child)
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
	for _, s := range ctx.AllStatement() {
		results = append(results, visitor.Visit(s))
	}
	return results
}

func (visitor *GogenVisitor) VisitStatement(ctx *parser.StatementContext) any {
	switch {
	case ctx.Statement(0) != nil:
		fmt.Printf("if paren_expr statement\n")
	case ctx.Statement(1) != nil:
		fmt.Printf("if paren_expr statement else statement\n")
	case ctx.Statement(2) != nil:
		fmt.Printf("while paren_expr statement\n")
	case ctx.Statement(3) != nil:
		fmt.Printf("do statement while paren_expr ;\n")
	case ctx.Statement(4) != nil:
		fmt.Printf("{ statement* }\n")
	case ctx.Statement(5) != nil:
		fmt.Printf("expr ;\n")
	case ctx.Statement(6) != nil:
		fmt.Printf(";\n")
	}
	return nil
}

func (visitor *GogenVisitor) IfNoElseStatement(stmt *parser.StatementContext) any {	
	s := strings.Builder{}
	s.WriteString("if")
	parenExprStr := visitor.Visit(stmt.Paren_expr()).(string)
	s.WriteString(parenExprStr)
	stmtStr := visitor.Visit(stmt.Statement(0)).(string)
	s.WriteString(stmtStr)
	return s.String()
}

func (visitor *GogenVisitor) VisitParen_expr(ctx *parser.Paren_exprContext) any {
	return nil
}

func (visitor *GogenVisitor) VisitExpr(ctx *parser.ExprContext) any {
	return nil 
}

func (visitor *GogenVisitor) VisitTest(ctx *parser.TestContext) any {
	switch {
	case ctx.Sum_(0) != nil:
		return visitor.Visit(ctx.Sum_(0))
	}	
	return nil
}

func (visitor *GogenVisitor) VisitAssignment(ctx *parser.ExprContext) any {
	s := strings.Builder{}
	s.WriteString(ctx.Id_().GetText())
	s.WriteString(" = ")
	s.WriteString(visitor.Visit(ctx.Expr()))
	return s.String()
}

func (visitor *GogenVisitor) VisitSum_(ctx *parser.Sum_Context) any {
	switch {
	case ctx.Term() != nil:
		return visitor.Visit(ctx.Term())
	case ctx.Sum_() != nil:
		return visitor.VisitBInarySum_(ctx.Sum_())
	}
	return nil
}

func (visitor *GogenVisitor) VisitBinarySum_(ctx *parser.Sum_Context) any {
	switch {
	case ctx.PLUS() != nil:
		return visitor.VisitSum_Add(ctx)	
	case ctx.MINUS() != nil:
		return visitor.VisitSum_Sub(ctx)
	}
	return nil
}

func (visitor *GogenVisitor) VisitSum_Add(ctx *parser.Sum_Context) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Sum_()))
	s.WriteString(" + ")
	s.WriteString(visitor.Visit(ctx.Term()))
	return s.String()
}	

func (visitor *GogenVisitor) VisitSum_Sub(ctx *parser.Sum_Context) any {
	s := strings.Builder{}
	s.WriteString(visitor.Visit(ctx.Sum_()))
	s.WriteString(" - ")
	s.WriteString(visitor.Visit(ctx.Term()))
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
	return nil
}

func (visitor *GogenVisitor) VisitInteger(ctx *parser.IntegerContext) any {
	return nil 
}

