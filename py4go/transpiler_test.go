package py4go

import (
	"fmt"
	"testing"
)

func TestTranspile(t *testing.T) {
	filenames := []string{
		"../testdata/p.py",
		"../testdata/PyFEM-master/pyfem/solvers/RiksSolver.py",
	}
	for index, f := range filenames {
		asttree, err := Parse(f)
		if err != nil {
			t.Fatal(err)
		}
		code, err := Transpile(asttree)
		if err != nil {
			t.Fatal(err)
		}
		testCase(t, fmt.Sprintf("../testdata/%d.go", index), code)
	}
}
