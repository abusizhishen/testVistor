// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Vistor

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseVistorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVistorVisitor) VisitNum(ctx *NumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitCalculate(ctx *CalculateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitElseIfStatement(ctx *ElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitElseStatement(ctx *ElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVistorVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}
