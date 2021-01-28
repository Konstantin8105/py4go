package py4go

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Parse python file to ast python tree
func Parse(filename string) (nodes Node, err error) {
	// minimal python code to ast tree generate
	code := `
import ast
r = open('%s','r')
t = ast.parse(r.read())
print(ast.dump(t))
`
	// current folder
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	filename = dir + string(os.PathSeparator) + filename
	fmt.Println(filename)

	// create transpiler file
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return
	}
	defer os.Remove(tmpfile.Name()) // clean up
	code = fmt.Sprintf(code, filename)
	if _, err = tmpfile.Write([]byte(code)); err != nil {
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
		err = fmt.Errorf("%v\n%v", err, func() error {
			cmd := exec.Command("python", tmpfile.Name())
			bs, err := cmd.CombinedOutput()
			return fmt.Errorf("%v\n%v\n%s", string(bs), err, out.String())
		}())
		return
	}

	// covert ast line string to nodes
	return ast(out.String())
}
