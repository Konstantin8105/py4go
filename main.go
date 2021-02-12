package main

import (
	"fmt"
	"os"

	"github.com/Konstantin8105/py4go/py4go"
)

func main() {
	for _, file := range os.Args[1:] {
		out, err := py4go.Parse(file)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	}
}
