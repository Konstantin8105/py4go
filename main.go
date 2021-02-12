package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Konstantin8105/py4go/py4go"
)

var (
	transpileCommand = flag.NewFlagSet(
		"transpile", flag.ContinueOnError)
	transpileHelpFlag = transpileCommand.Bool(
		"h", false, "print help information")

	astCommand = flag.NewFlagSet(
		"ast", flag.ContinueOnError)
	astHelpFlag = astCommand.Bool(
		"h", false, "print help information")

		stderr = os.Stderr
)

func main() {

	flag.Usage = func() {
		usage := "Usage: %s [<command>] [<flags>] file1.py ...\n\n"
		usage += "Commands:\n"
		usage += "  transpile\ttranspile an input Python source file or files to Go\n"
		usage += "  ast\t\tprint AST before translated Go code\n"
		usage += "\n"
		fmt.Fprintf(stderr, usage, os.Args[0])

		flag.PrintDefaults()
	}

	transpileCommand.SetOutput(stderr)
	astCommand.SetOutput(stderr)

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	var isAst, isTranspile bool

	switch os.Args[1] {
	case "ast":
		err := astCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stdout, "ast command cannot parse: %v", err)
			return
		}

		if *astHelpFlag || astCommand.NArg() == 0 {
			fmt.Fprintf(stderr, "Usage: %s ast file.py ...\n", os.Args[0])
			astCommand.PrintDefaults()
			return
		}

		isAst = true

	case "transpile":
		err := transpileCommand.Parse(os.Args[2:])
		if err != nil {
			fmt.Fprintf(os.Stdout, "transpile command cannot parse: %v", err)
			return
		}

		if *transpileHelpFlag || transpileCommand.NArg() == 0 {
			fmt.Fprintf(stderr,
				"Usage: %s transpile file1.py ...\n",
				os.Args[0])
			transpileCommand.PrintDefaults()
			return
		}

		isTranspile = true

	default:
		flag.Usage()
		return
	}

	for _, file := range os.Args[2:] {
		asttree, err := py4go.Parse(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if isAst {
			fmt.Println(asttree)
			continue
		}
		if !isTranspile {
			continue
		}
		code, err := py4go.Transpile(asttree)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(code)
	}
}
