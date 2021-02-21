// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Vistor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VistorListener is a complete listener for a parse tree produced by VistorParser.
type VistorListener interface {
	antlr.ParseTreeListener

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterCalculate is called when entering the calculate production.
	EnterCalculate(c *CalculateContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterElseIfStatement is called when entering the elseIfStatement production.
	EnterElseIfStatement(c *ElseIfStatementContext)

	// EnterElseStatement is called when entering the elseStatement production.
	EnterElseStatement(c *ElseStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitCalculate is called when exiting the calculate production.
	ExitCalculate(c *CalculateContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitElseIfStatement is called when exiting the elseIfStatement production.
	ExitElseIfStatement(c *ElseIfStatementContext)

	// ExitElseStatement is called when exiting the elseStatement production.
	ExitElseStatement(c *ElseStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)
}
