package py4go

import "testing"

func TestTranspile(t *testing.T) {
	filename := "../testdata/p.py"
	asttree, err := Parse(filename)
	if err != nil {
		t.Fatal(err)
	}
	code, err := Transpile(asttree)
	if err != nil {
		t.Fatal(err)
	}
	testCase(t, "../testdata/.p.py.go", code)
}
