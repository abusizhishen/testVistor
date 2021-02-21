// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Vistor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by VistorParser.
type VistorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VistorParser#num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by VistorParser#calculate.
	VisitCalculate(ctx *CalculateContext) interface{}

	// Visit a parse tree produced by VistorParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by VistorParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by VistorParser#elseIfStatement.
	VisitElseIfStatement(ctx *ElseIfStatementContext) interface{}

	// Visit a parse tree produced by VistorParser#elseStatement.
	VisitElseStatement(ctx *ElseStatementContext) interface{}

	// Visit a parse tree produced by VistorParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by VistorParser#start.
	VisitStart(ctx *StartContext) interface{}
}
