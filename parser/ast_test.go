package parser

import (
	"fmt"
	"testing"
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

func TestAst(t *testing.T) {
	tcs := []string{
		"arguments(args)",
		"arguments(args())",
		"arguments(args,vararg)",
		"arguments(args[],vararg)",
		"arguments(vararg,args[])",
		"arguments(vararg,args[],vararg)",
		"arguments(args[],vararg,args[])",
		"(args[],vararg,args[])",
		"vararg=None",
		"(vararg=None)",
		"arguments(vararg=None)",
		"arguments(args[],vararg=None)",
		"arguments(vararg=None,args[])",
		"arguments(args[],vararg=None,kwarg=None,defaults[])",
		"body[n=3.1415]",
		"body[Return(value=Num(n=3.1415))]",
		"decorator_list[]",
		"FunctionDef(name='print_x', args=arguments(args[], vararg=None, kwarg=None, defaults[]))",
		func() string {
			filename := "../testdata/p.py"
			out, err := Parse(filename)
			if err != nil {
				t.Fatal(err)
			}
			return out
		}(),
	}

	for i := range tcs {
		t.Run(fmt.Sprintf("%2dT%s", i, tcs[i]), func(t *testing.T) {
			a, err := Ast(tcs[i])
			if err != nil {
				t.Fatal(err)
			}
			t.Logf("\n%v", a)
		})
	}
}
