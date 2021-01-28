package py4go

import (
	"fmt"
	"os"
)

func ExampleTranspile() {
	filename := "../testdata/p.py"
	asttree, err := Parse(filename)
	if err != nil {
		panic(err)
	}
	// TODO: fmt.Fprintln(os.Stdout, asttree)
	code, err := Transpile(asttree)
	fmt.Fprintln(os.Stdout, code, "\n", err)
	// Output:
}
