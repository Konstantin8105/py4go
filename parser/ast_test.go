package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
	}

	for i := range tcs {
		t.Run(fmt.Sprintf("%2dT%s", i, tcs[i]), func(t *testing.T) {
			a, err := ast(tcs[i])
			if err != nil {
				t.Fatal(err)
			}
			testCase(t, fmt.Sprintf("../testdata/.ast%02d", i), a.String())
		})
	}
}

func TestIntegration(t *testing.T) {
	// FULL=true go test
	if os.Getenv("FULL") != "true" {
		return
	}

	err := filepath.Walk("../testdata/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".py") {
			t.Run(path, func(t *testing.T) {
				_, err := Parse(path)
				if err != nil {
					return
				}
			})

		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
