package parser

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Parse(filename string) (_ string, err error) {
	// minimal python code to ast tree generate
	code := `
import ast
r = open('%s','r')
t = ast.parse(r.read())
print(ast.dump(t))
`

	// create transpiler file
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return
	}
	defer os.Remove(tmpfile.Name()) // clean up
	if _, err = tmpfile.Write([]byte(fmt.Sprintf(code, filename))); err != nil {
		return
	}
	if err = tmpfile.Close(); err != nil {
		return
	}

	// get ast tree
	cmd := exec.Command("python", tmpfile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		return
	}

	return out.String(), nil
}
