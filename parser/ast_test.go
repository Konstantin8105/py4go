package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

func testCase(t *testing.T, filename, result string) {
	t.Run(filename, func(t *testing.T) {
		// for update test screens run in console:
		// UPDATE=true go test
		if os.Getenv("UPDATE") == "true" {
			if err := ioutil.WriteFile(filename, []byte(result), 0644); err != nil {
				t.Fatalf("Cannot write snapshot to file: %v", err)
			}
		}

		// compare datas
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			t.Fatalf("Cannot read snapshot file: %v", err)
		}
		if !bytes.Equal([]byte(result), content) {
			t.Errorf("Snapshots is not same:\n%s\n%s", result, string(content))
		}
	})
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
		"(value=Num(n=3.1415), a)",
		"(value=Num(n=3.1415), a[])",
		"Assign(targets=[Name(id='x', ctx=Store())], value=Num(n=1))",
		"body[Return(value=Num(n=3.1415))]",
		"b(a(c(),d),e)",
		"b=(a(c,d()),e)",
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
			testCase(t, fmt.Sprintf("../testdata/.ast%02d", i), a.String())
		})
	}
}
