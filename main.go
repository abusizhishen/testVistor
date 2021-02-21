package main

import (
	"fmt"
	"github.com/abusizhishen/testVistor/parser"
	"github.com/abusizhishen/testVistor/src"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input,err := antlr.NewFileStream("test.ru")
	if err != nil{
		panic(err)
	}

	lexer := parser.NewVistorLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewVistorParser(tokens)

	v := new(src.BaseVisitor)
	value := p.Start().Accept(v)
	fmt.Println(value)

}
