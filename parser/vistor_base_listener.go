// Code generated from Vistor.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Vistor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVistorListener is a complete listener for a parse tree produced by VistorParser.
type BaseVistorListener struct{}

var _ VistorListener = &BaseVistorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVistorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVistorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVistorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVistorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNum is called when production num is entered.
func (s *BaseVistorListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseVistorListener) ExitNum(ctx *NumContext) {}

// EnterCalculate is called when production calculate is entered.
func (s *BaseVistorListener) EnterCalculate(ctx *CalculateContext) {}

// ExitCalculate is called when production calculate is exited.
func (s *BaseVistorListener) ExitCalculate(ctx *CalculateContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseVistorListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseVistorListener) ExitCondition(ctx *ConditionContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseVistorListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseVistorListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterElseIfStatement is called when production elseIfStatement is entered.
func (s *BaseVistorListener) EnterElseIfStatement(ctx *ElseIfStatementContext) {}

// ExitElseIfStatement is called when production elseIfStatement is exited.
func (s *BaseVistorListener) ExitElseIfStatement(ctx *ElseIfStatementContext) {}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseVistorListener) EnterElseStatement(ctx *ElseStatementContext) {}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseVistorListener) ExitElseStatement(ctx *ElseStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseVistorListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseVistorListener) ExitStatement(ctx *StatementContext) {}

// EnterStart is called when production start is entered.
func (s *BaseVistorListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseVistorListener) ExitStart(ctx *StartContext) {}
