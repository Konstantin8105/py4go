package py4go_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Konstantin8105/py4go/py4go"
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
				_, err := py4go.Parse(path)
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
