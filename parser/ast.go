package parser

import (
	"container/list"
	"fmt"
	"go/scanner"
	"go/token"
	"strings"
)

type Node struct {
	Name  string
	Nodes []*Node
}

func (n *Node) RemoveEmpty() {
	for i := 0; i < len(n.Nodes); i++ {
		n.Nodes[i].RemoveEmpty()
	}
	for i := 0; i < len(n.Nodes); i++ {
		if v := n.Nodes[i]; v.Name == "" && len(v.Nodes) == 0 {
			n.Nodes = append(n.Nodes[:i], n.Nodes[i+1:]...)
			n.RemoveEmpty()
			return
		}
	}
}

func (n Node) String() string {
	var str string
	if n.Name != "" {
		str += n.Name
	} else {
		str += "NONAME"
	}
	str += "(\n"
	for i := range n.Nodes {
		ss := strings.Split(n.Nodes[i].String(), "\n")
		for j := range ss {
			str += " " + ss[j] + "\n"
		}
	}
	str += ") // " + n.Name
	return str
}

func Ast(src string) (nodes *Node, err error) {

	src = strings.ReplaceAll(src, "=[", "[")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                               // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src))          // register input "file"
	s.Init(file, []byte(src), nil /* no error handler */, 0) // scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	l := list.New()

	for {
		pos, tok, lit := s.Scan()
		_ = pos
		if tok == token.EOF {
			break
		}
		if tok == token.SEMICOLON {
			continue
		}

		s := lit
		if s == "" {
			s = fmt.Sprintf("%s", tok)
		}

		if nodes == (*Node)(nil) {
			nodes = new(Node)
			nodes.Name = s
			l.PushFront(nodes)
			continue
		}

		if tok == token.LPAREN ||
			tok == token.ASSIGN ||
			tok == token.LBRACK {
			// add new nodes in last
			e := l.Front().Value.(*Node)
			e.Nodes = append(e.Nodes, new(Node))
			l.PushFront(e.Nodes[len(e.Nodes)-1])
			continue
		}
		if tok == token.RPAREN || tok == token.RBRACK {
			// going outside
			l.Remove(l.Front())
			continue
		}
		if tok == token.COMMA {
			l.Remove(l.Front())
			l.Remove(l.Front())
			e := l.Front().Value.(*Node)
			e.Nodes = append(e.Nodes, new(Node))
			l.PushFront(e.Nodes[len(e.Nodes)-1])
			continue
		}
		l.Front().Value.(*Node).Name = s
	}

	if len(nodes.Nodes) != 1 {
		panic(fmt.Errorf("strange : %v", len(nodes.Nodes)))
	}

	// remove empty nodes
	nodes.RemoveEmpty()

	return
}
