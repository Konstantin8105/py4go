package py4go

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

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
				ast, err := Parse(path)
				if err != nil {
					return
				}
				code, err := Transpile(ast)
				t.Logf("%s", ast)
				t.Logf("%s", code)
				if err != nil {
					t.Errorf("%s", err)
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
