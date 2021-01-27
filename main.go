package main

import (
	"fmt"

	"github.com/Konstantin8105/py4go/py4go"
)

func main() {
	filename := "testdata/p.py"
	out, err := py4go.Parse(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
