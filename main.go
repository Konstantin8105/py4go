package main

import (
	"fmt"

	"github.com/Konstantin8105/py4go/parser"
)

func main() {
	filename := "/home/konstantin/go/src/github.com/Konstantin8105/py4go/p.py"
	out, err := parser.Parse(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	ast, err := parser.Ast(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(ast)
}
