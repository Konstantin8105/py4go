package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	filename := "/home/konstantin/go/src/github.com/Konstantin8105/py4go/p.py"

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
		panic(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up
	if _, err := tmpfile.Write([]byte(fmt.Sprintf(code, filename))); err != nil {
		panic(err)
	}
	if err := tmpfile.Close(); err != nil {
		panic(err)
	}

	// get ast tree
	cmd := exec.Command("python", tmpfile.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// print ast tree
	fmt.Println(out.String())
}
