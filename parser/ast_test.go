package parser

import (
	"fmt"
)

func ExampleAst() {
	str := "Module(body=[Assign(targets=[Name(id='x', ctx=Store())], value=Num(n=1))])"
	n, err := Ast(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
	fmt.Println(n.String())
	// Output:
	// Module(body=[Assign(targets=[Name(id='x', ctx=Store())], value=Num(n=1))])
	// Module(
	//  body(
	//   Assign(
	//    targets(
	//     Name(
	//      id(
	//       'x'(
	//       ) // 'x'
	//      ) // id
	//      ctx(
	//       Store(
	//       ) // Store
	//      ) // ctx
	//     ) // Name
	//    ) // targets
	//    value(
	//     Num(
	//      n(
	//       1(
	//       ) // 1
	//      ) // n
	//     ) // Num
	//    ) // value
	//   ) // Assign
	//  ) // body
	// ) // Module
}
