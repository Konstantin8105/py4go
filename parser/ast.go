package parser

import (
	"container/list"
	"fmt"
	"go/scanner"
	"go/token"
	"strings"
)

type Node interface {
	String() string
}

type Ident struct {
	Name string
}

func (id Ident) String() string {
	return id.Name
}

type List struct {
	IsParen bool // false - [], true - ()
	Name    string
	Args    []Node
}

func (l List) String() string {
	var str string
	for _, v := range l.Args {
		lines := strings.Split(v.String(), "\n")
		for _, line := range lines {
			str += fmt.Sprintf("  %s\n", line)
		}
	}
	if l.IsParen {
		return fmt.Sprintf("%s (\n%s) // %s", l.Name, str, l.Name)
	}
	return fmt.Sprintf("%s [\n%s] // %s", l.Name, str, l.Name)
}

type Assign struct {
	Left, Right Node
}

func (a Assign) String() string {
	return fmt.Sprintf("%v = %v", a.Left, a.Right)
}

func Ast(src string) (nodes Node, err error) {

	// scan
	var (
		// initialize the scanner.
		s scanner.Scanner
		// positions are relative to fset
		fset = token.NewFileSet()
		// register input "file"
		file = fset.AddFile("", fset.Base(), len(src))
	)
	s.Init(file, []byte(src), nil, 0)

	// store all tokens
	type Element struct {
		tok token.Token
		str string
	}
	elements := []Element{}
	for {
		_, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		str := lit
		if str == "" {
			str = fmt.Sprintf("%s", tok)
		}
		elements = append(elements, Element{tok: tok, str: str})
	}

	// convert to nodes
	l := list.New()
	for i := 0; i < len(elements); i++ {

		// fmt.Println(">>>>>>>>>>>>>", elements[:i])
		// for e := l.Front(); e != nil; e = e.Next() {
		// 	fmt.Printf("%p %#v\n", e.Value, e.Value)
		// }
		// fmt.Println("Nodes: ", nodes)
		// fmt.Println("TOK : ", elements[i].tok)
		// fmt.Println("STR : ", elements[i].str)
		// fmt.Println("LEN : ", l.Len())

		switch elements[i].tok {
		case token.SEMICOLON:
			continue
		case token.ASSIGN:
			// example:
			// A   = ...
			// A() = ...
			// A[] = ...
			fr := l.Front().Value.(Node)
			switch fr.(type) {
			case (*List), (*Ident):
				d := l.Remove(l.Front()).(Node) // Ident or List
				a := Assign{Left: d}            // Right element later
				a.Right = &Ident{}
				if l.Len() == 0 {
					l.PushFront(&a)
				} else {
					if list, ok := l.Front().Value.(*List); ok {
						list.Args[len(list.Args)-1] = &a
					} else {
						panic(l)
					}
				}
				l.PushFront(a.Right)
			default:
				panic(fr)
			}

		case token.COMMA:
			for {
				l.Remove(l.Front())
				if _, ok := l.Front().Value.(*List); ok {
					break
				}
			}
			fr := l.Front().Value.(Node)
			switch fr.(type) {
			case (*List):
				list := fr.(*List)
				id := Ident{}
				list.Args = append(list.Args, &id)
				l.PushFront(&id)
			default:
				panic(fr)
			}

		case token.LPAREN, token.LBRACK:
			// if before ident, then it if named list
			var fr Node
			if 0 < l.Len() {
				fr = l.Front().Value.(Node)
			} else {
				var list List
				l.PushFront(&list)
				continue
			}
			list := List{IsParen: elements[i].tok == token.LPAREN}
			list.Args = append(list.Args, new(Ident))
			switch fr.(type) {
			case (*Ident):
				// example: A(...
				id := fr.(*Ident)
				list.Name = id.Name
				l.Remove(l.Front())
				if l.Len() == 0 {
					l.PushFront(&list)
				} else {
					if ll, ok := l.Front().Value.(*List); ok {
						if a2, ok := ll.Args[len(ll.Args)-1].(*Assign); ok {
							a2.Right = &list
						} else {
							ll.Args[len(ll.Args)-1] = &list
						}
						l.PushFront(&list)
					} else if a, ok := l.Front().Value.(*Assign); ok {
						a.Right = &list
						l.PushFront(&list)
					} else {
						panic(l.Front().Value)
					}
				}
			default:
				// example: ... = (...)
				l.PushFront(&list)
			}
			l.PushFront(list.Args[0])

		case token.RPAREN, token.RBRACK:
			l.Remove(l.Front())
			fr := l.Front().Value.(Node)
			if _, ok := fr.(*List); !ok {
				l.Remove(l.Front())
			}

		default:
			// Ident
			id := Ident{Name: elements[i].str}
			if l.Front() == nil {
				l.PushFront(&id)
				continue
			}
			fr := l.Front().Value.(Node)
			switch fr.(type) {
			case *List:
				list := fr.(*List)
				list.Args = append(list.Args, &id)
				l.PushFront(&id)
			case *Assign:
				a := fr.(*Assign)
				a.Right = &id
				l.PushFront(&id)
			case *Ident:
				last := fr.(*Ident)
				last.Name = id.Name
			default:
				panic(fr)
			}
		}
		nodes = l.Back().Value.(Node)
	}

	return
}
