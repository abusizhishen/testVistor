package src

import (
	"fmt"
	"github.com/abusizhishen/testVistor/parser"
	"strconv"
)

type BaseVisitor struct {
	*parser.BaseVistorVisitor
	stack []int
}

func (v *BaseVisitor)pop() int {
	if len(v.stack) < 1 {
		panic("stack is empty")
	}
	num := v.stack[len(v.stack)-1]
	v.stack = v.stack[:len(v.stack)-1]

	return num
}

func (v *BaseVisitor)push(num int)  {
	v.stack = append(v.stack,num)
}

func (v *BaseVisitor)VisitCalculate(ctx *parser.CalculateContext) interface{} {
	fmt.Println("VisitCalculate", ctx.GetText())
	ctx.Num(0).Accept(v)
	ctx.Num(1).Accept(v)
	right,left := v.pop(),v.pop()

	switch ctx.GetOp().GetText() {
	case "+":
		v.push(left+right)
	case "-":
		v.push(left-right)
	case "*":
		v.push(left*right)
	case "/":
		if right == 0{
			panic("除数不能为0")
		}
		v.push(left/right)
	}

	return nil
}

func (v *BaseVisitor)VisitNum(ctx *parser.NumContext) interface{} {
	fmt.Println("VisitNum", ctx.GetText())

	s := ctx.NUM().GetText()
	num,err := strconv.Atoi(s)
	if err != nil{
		panic(err)
	}

	v.push(num)
	return nil
}

func (v *BaseVisitor)VisitStart(ctx *parser.StartContext) interface{} {
	fmt.Println("VisitStart", ctx.GetText())
	for _,statement := range ctx.AllStatement(){
		statement.Accept(v)
	}

	fmt.Println("stack",v.stack)
	return v.pop()
}

func (v *BaseVisitor)VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	fmt.Println("VisitIfStatement", ctx.GetText())
	if ctx.Condition().Accept(v).(bool){
		for _,statement := range ctx.AllStatement(){
			statement.Accept(v)
		}
		return nil
	}

	if len(ctx.AllElseIfStatement())==0{
		ctx.ElseStatement()
		return nil
	}

	for _,statement := range ctx.AllElseIfStatement(){
		if statement.Accept(v).(bool) {
			return nil
		}
	}

	ctx.ElseStatement().Accept(v)
	return nil
}

func (v *BaseVisitor)VisitCondition(ctx *parser.ConditionContext) interface{} {
	fmt.Println("VisitCondition", ctx.GetText())
	switch ctx.GetText() {
	case "true":
		return true
	case "false":
		return false
	default:
		panic("无效的"+ctx.GetText())
	}
}

func (v *BaseVisitor)VisitElseIfStatement(ctx *parser.ElseIfStatementContext) interface{} {
	fmt.Println("VisitElseIfStatement", ctx.GetText())

	if !ctx.Condition().Accept(v).(bool){
		return false
	}

	for _,statement := range ctx.AllStatement(){
		statement.Accept(v)
	}

	return true
}

func (v *BaseVisitor)VisitElseStatement(ctx *parser.ElseStatementContext) interface{} {
	fmt.Println("VisitElseStatement", ctx.GetText())

	for _,statement := range ctx.AllStatement(){
		statement.Accept(v)
	}

	return nil
}

func (v *BaseVisitor)VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("VisitStatement", ctx.GetText())

	for i:=0;i<ctx.GetChildCount();i++{
		ctx.GetChildOfType(i,nil).Accept(v)
	}

	return nil
}