package gogen

import (
	"errors"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/threadedstream/tinycgo/internal/pkg/aid"
	"github.com/threadedstream/tinycgo/internal/pkg/logger"
	"github.com/threadedstream/tinycgo/internal/pkg/scope"
	parser "github.com/threadedstream/tinycgo/pkg/antlr"
	"reflect"
	"strconv"
)

// EvalVisitor is a visitor designed mainly for evaluation of condition expressions
type EvalVisitor struct {
	scope *scope.Scope
}

// NewEvalVisitor returns a fresh instance of EvalVisitor
func NewEvalVisitor() *EvalVisitor {
	return &EvalVisitor{}
}

func (ev *EvalVisitor) Visit(root antlr.Tree) any {
	switch child := root.(type) {
	case *parser.ExprContext:
		return ev.VisitExpr(child)
	case *parser.AssignmentExprContext:
		return ev.VisitAssignmentExpr(child)
	case *parser.TestContext:
		return ev.VisitTest(child)
	case *parser.Sum_Context:
		return ev.VisitSum_(child)
	case *parser.CompareSum_Context:
		return ev.VisitCompareSum_(child)
	case *parser.TermContext:
		return ev.VisitTerm(child)
	}
	return nil
}

func (ev *EvalVisitor) VisitExpr(ctx *parser.ExprContext) any {
	switch {
	case ctx.Test() != nil:
		return ev.Visit(ctx.Test())
	case ctx.AssignmentExpr() != nil:
		return ev.Visit(ctx.AssignmentExpr())
	}
	return nil
}

func (ev *EvalVisitor) VisitAssignmentExpr(_ *parser.AssignmentExprContext) bool {
	return true
}

func (ev *EvalVisitor) VisitTest(ctx *parser.TestContext) bool {
	switch {
	case ctx.Sum_() != nil:
		var err error
		var val bool
		raw := ev.VisitSum_(ctx.Sum_())
		if val, err = ev.coalesceToBool(raw); err == nil {
			return val
		}
		logger.Error(err.Error())
	case ctx.CompareSum_() != nil:
		return ev.VisitCompareSum_(ctx.CompareSum_())
	}
	return false
}

func (ev *EvalVisitor) VisitSum_(ctx parser.ISum_Context) any {
	sumCtx := ctx.(*parser.Sum_Context)
	switch {
	case sumCtx.Term() != nil:
		return ev.VisitTerm(sumCtx.Term())
	case sumCtx.Sum_() != nil:
		return ev.VisitBinaryExpression(sumCtx)
	}
	return false
}

func (ev *EvalVisitor) VisitBinaryExpression(ctx *parser.Sum_Context) int {
	var normedLhs, normedRhs int
	var err error
	lhs := ev.Visit(ctx.Sum_())
	rhs := ev.Visit(ctx.Term())
	if normedLhs, err = ev.coalesceToInt(lhs); err != nil {
		logger.Error(err.Error())
	}
	if normedRhs, err = ev.coalesceToInt(rhs); err != nil {
		logger.Error(err.Error())
	}
	binopCtx := ctx.Binop().(*parser.BinopContext)
	switch {
	default:
		logger.Error("expected binop to be '+' or '-' ")
	case binopCtx.PLUS() != nil:
		return normedLhs + normedRhs
	case binopCtx.MINUS() != nil:
		return normedLhs - normedRhs
	}
	// not supposed to reach that
	return -69
}

func (ev *EvalVisitor) VisitCompareSum_(ctx parser.ICompareSum_Context) bool {
	compareSumCtx := ctx.(*parser.CompareSum_Context)
	var normedLhs, normedRhs int
	var err error
	lhs := ev.Visit(compareSumCtx.Sum_(0))
	rhs := ev.Visit(compareSumCtx.Sum_(1))
	if normedLhs, err = ev.coalesceToInt(lhs); err != nil {
		logger.Error(err.Error())
	}
	if normedRhs, err = ev.coalesceToInt(rhs); err != nil {
		logger.Error(err.Error())
	}
	return normedLhs < normedRhs
}

func (ev *EvalVisitor) VisitTerm(ctx parser.ITermContext) any {
	termCtx := ctx.(*parser.TermContext)
	switch {
	case termCtx.Id_() != nil:
		if raw, found := ev.scope.Get(termCtx.Id_().GetText()); found {
			if val, isInt := raw.(int); isInt {
				return val
			}
			logger.Error("value must be of type integer")
		}
		logger.Error("no entry with name %s", termCtx.Id_().GetText())
	case termCtx.Integer() != nil:
		if cvt, err := strconv.ParseInt(termCtx.Integer().GetText(), 10, 64); err == nil {
			// cvt is of type int64
			val := int(cvt)
			return val
		} else {
			logger.Error("failed to convert string to integer, reason: %s", err.Error())
		}
	case termCtx.Paren_expr() != nil:
		res := ev.Visit(termCtx.Paren_expr())
		ev.assertOneOf(res, reflect.Bool, reflect.Int)
	}
	return nil
}

func (ev *EvalVisitor) assertOneOf(raw any, tys ...reflect.Kind) {
	ty := reflect.ValueOf(raw)
	if aid.Contains(ty, tys) {
		return
	} else {
		logger.Error("raw's type is not in tys")
	}
}

func (ev *EvalVisitor) coalesceToInt(raw any) (int, error) {
	if val, isInt := raw.(int); isInt {
		return val, nil
	}
	if val, isBool := raw.(bool); isBool {
		return ev.coalesceBoolToInt(val), nil
	}
	return -1, errors.New("expected raw to be of type int or bool")
}

func (ev *EvalVisitor) coalesceToBool(raw any) (bool, error) {
	if val, isInt := raw.(int); isInt {
		return ev.coalesceIntToBool(val), nil
	}
	if val, isBool := raw.(bool); isBool {
		return val, nil
	}
	return false, errors.New("expected raw to be of type int or bool")
}

func (ev *EvalVisitor) coalesceBoolToInt(val bool) int {
	if val {
		return 1
	} else {
		return 0
	}
}

func (ev *EvalVisitor) coalesceIntToBool(val int) bool {
	if val == 0 {
		return false
	}
	return true
}
