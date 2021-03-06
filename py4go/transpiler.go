package py4go

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"go/format"
	"go/printer"
	"go/token"
	"runtime/debug"
	"strings"

	"github.com/Konstantin8105/errors"
)

func Transpile(nodes Node) (gocode string, err error) {
	gonode := &goast.File{
		Name: goast.NewIdent("main"),
	}

	// tranpiling
	decls, stmts, errDecls := transpile(nodes)
	defer func() {
		if errDecls != nil {
			err = fmt.Errorf("%v\n%v", err, errDecls)
		}
	}()

	if 0 < len(stmts) {
		decls = append(decls, &goast.FuncDecl{
			Name: goast.NewIdent("main"),
			Type: &goast.FuncType{
				Params: &goast.FieldList{},
			},
			Body: &goast.BlockStmt{List: stmts},
		})
	}

	gonode.Decls = decls

	// Create a FileSet for node. Since the node does not come
	// from a real source file, fset will be empty.
	fset := token.NewFileSet()

	var buf bytes.Buffer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(nodes)
			_ = goast.Print(fset, gonode)
			err = fmt.Errorf(
				"func Transpile\nstacktrace from panic: \n %s\n %v",
				string(debug.Stack()),
				err,
			)
		}
	}()
	err = format.Node(&buf, fset, gonode)
	if err != nil {
		_ = goast.Print(fset, gonode)
		return
	}
	return buf.String(), err
}

func view(g interface{}) {
	fset := token.NewFileSet()
	_ = goast.Print(fset, g)

	var buf bytes.Buffer
	_ = format.Node(&buf, token.NewFileSet(), g)

	fmt.Println(buf.String())
}

func clearName(str string) string {
	return strings.Replace(str, "'", "", -1)
}

func isIdent(n Node, name string) (ind *Ident, ok bool) {
	ind, ok = n.(*Ident)
	if !ok {
		return nil, false
	}
	if ind.Name != name {
		return nil, false
	}
	return ind, true
}

func isAssign(n Node, name string) (a *Assign, ok bool) {
	a, ok = n.(*Assign)
	if !ok {
		return nil, false
	}
	_, ok = isIdent(a.Left, name)
	if !ok {
		return nil, false
	}
	return a, true
}

func isList(n Node, name string) (list *List, ok bool) {
	list, ok = n.(*List)
	if !ok {
		return nil, false
	}
	if list.Name != name {
		return nil, false
	}
	return list, true
}

func exprsSel(exprs []goast.Expr) (e goast.Expr, err error) {
	if len(exprs) == 0 {
		err = fmt.Errorf("right is zero")
		return
	}
	e = exprs[0]
	for k := 1; k < len(exprs); k++ {
		if id, ok := exprs[k].(*goast.Ident); ok {
			e = &goast.SelectorExpr{
				X:   e,
				Sel: id,
			}
		} else {
			e = &goast.IndexExpr{
				X:     e,
				Index: exprs[k],
			}
		}
	}
	return
}

func transpile(n Node) (decls []goast.Decl, stmts []goast.Stmt, err error) {
	et := errors.New("Error in func transpile")

	addStmt := func(n Node) {
		stmt, errStmt := transpileStmt(n)
		if errStmt != nil {
			et.Add(errStmt)
			return
		}
		if stmt != nil {
			stmts = append(stmts, stmt)
		}
	}

	switch v := n.(type) {
	case *List:
		switch v.Name {
		case "Module":
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "body"); ok {
						if right, ok := a.Right.(*List); ok {
							for j := range right.Args {
								ds, dstmts, err := transpile(right.Args[j])
								decls = append(decls, ds...)
								stmts = append(stmts, dstmts...)
								et.Add(err)
							}
						} else {
							et.Add(fmt.Errorf("expect Assign.Ident.List: %s", n))
						}
					} else {
						et.Add(fmt.Errorf("expect Assign.Ident: %s", n))
					}
				} else {
					et.Add(fmt.Errorf("expect Assign: %s", n))
				}
			}
		case "ClassDef":
			// ClassDef (
			//  name = 'RiksSolver'
			//  bases =  [
			//    Name (
			//      id = 'BaseModule'
			//    ) // Name
			//  ] //
			//  body =  [ ... ]
			// ) // ClassDef

			// TODO - add as part of struct
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "body"); ok {
						if l, ok := a.Right.(*List); ok {
							for k := range l.Args {
								d, s, e := transpile(l.Args[k])
								if e != nil {
									et.Add(e)
								} else {
									decls = append(decls, d...)
									stmts = append(stmts, s...)
								}
							}
						}
					}
				}
			}

		case "FunctionDef":
			decl := goast.FuncDecl{
				Type: &goast.FuncType{
					Params: &goast.FieldList{},
				},
				Body: &goast.BlockStmt{},
			}
			for i := range v.Args {
				a, ok := v.Args[i].(*Assign)
				if !ok {
					continue
				}
				if _, ok := isIdent(a.Left, "name"); ok {
					name := fmt.Sprintf("%s", a.Right)
					decl.Name = goast.NewIdent(clearName(name))

				}
				if _, ok := isIdent(a.Left, "body"); ok {
					var (
						errBody error
						body    []goast.Stmt
					)
					body, errBody = transpileStmts(a.Right)
					if errBody != nil {
						et.Add(errBody)
						continue
					}
					decl.Body.List = append(decl.Body.List, body...)
				}
			}
			decls = append(decls, &decl)
		default:
			addStmt(n)
		}
	default:
		addStmt(n)
	}
	if et.IsError() {
		err = et
		return
	}
	return
}

func transpileStmts(n Node) (stmts []goast.Stmt, err error) {
	et := errors.New("Error in func transpileStmts")
	switch v := n.(type) {
	case *List:
		for i := range v.Args {
			stmt, errStmt := transpileStmt(v.Args[i])
			if errStmt != nil {
				lines := strings.Split(fmt.Sprintf("%v", errStmt), "\n")
				stmts = append(stmts, &goast.ExprStmt{
					X: goast.NewIdent("/" + "/ Find PY4GO error"),
				})
				for k := range lines {
					stmts = append(stmts, &goast.ExprStmt{
						X: goast.NewIdent("/" + "/ " + lines[k]),
					})
				}
				continue
			}
			if stmt != nil {
				stmts = append(stmts, stmt)
			}
		}

	default:
		et.Add(fmt.Errorf("cannot transpile : %s", n))
	}
	if et.IsError() {
		err = et
		return
	}
	return
}

func transpileStmt(n Node) (stmt goast.Stmt, err error) {
	// 	defer func() {
	// 		view(stmt)
	// 	}()
	et := errors.New("Error in func transpileStmt")
	switch v := n.(type) {
	case *Ident:
		if v.Name == "" {
			// ignore empty ident
			break
		}
	case *List:
		switch v.Name {
		case "Import":
			// Import (
			//   names =  [
			//     alias (
			//       name = 'setuptools'
			//       asname = None
			//     ) // alias
			//   ] //
			// ) // Import
			break // ignore
		case "ImportFrom":
			// ImportFrom (
			//   module = 'pyfem.elements.Spring'
			//   names =  [
			//     alias (
			//       name = 'Spring'
			//       asname = None
			//     ) // alias
			//   ] //
			//   level = 0
			// ) // ImportFrom
			break // ignore

		case "AugAssign":
			// from:
			//
			// AugAssign (
			//   target = ...
			//   op = Mult ()
			//   value  = ...
			// )
			//
			// to:
			//
			// Assign (
			//   targets =  [
			//     <<target>>
			//   ] //
			//   value  = BinOp(
			//       left  = <<target>>
			//       op    = <<op>>
			//       right = <<value>>
			//   )
			// )
			var (
				target Node
				op     Node
				value  Node
			)
			for i := range v.Args {
				if a, ok := isAssign(v.Args[i], "target"); ok {
					target = a.Right
				}
				if a, ok := isAssign(v.Args[i], "op"); ok {
					op = a.Right
				}
				if a, ok := isAssign(v.Args[i], "value"); ok {
					value = a.Right
				}
			}
			newassign := &List{
				Name: "Assign",
				Args: []Node{
					&Assign{
						Left:  &Ident{Name: "targets"},
						Right: &List{Args: []Node{target}},
					},
					&Assign{
						Left: &Ident{Name: "value"},
						Right: &List{
							Name: "BinOp",
							Args: []Node{
								&Assign{Left: &Ident{Name: "left"}, Right: target},
								&Assign{Left: &Ident{Name: "op"}, Right: op},
								&Assign{Left: &Ident{Name: "right"}, Right: value},
							},
						},
					},
				},
			}
			stmt, err = transpileStmt(newassign)
			if err != nil {
				et.Add(err)
			}

		case "Raise":
			// Raise (
			//   type = Call ( ... )
			// ) // Raise
			call := v.Args[0].(*Assign).Right
			exprsC, errC := transpileExprs(call)
			if errC != nil {
				et.Add(errC)
			} else {
				s, errs := exprsSel(exprsC)
				if errs != nil {
					et.Add(errs)
				} else {
					stmt = &goast.ExprStmt{
						X: s,
					}
				}
			}

		case "Assign":
			// from:
			//
			// Assign (
			//   targets =  [
			//     Name (
			//       id = 'x'
			//     ) // Name
			//   ] //
			//   value = Num (
			//     n = 0
			//   ) // Num
			// ) // Assign
			//
			// to:
			//
			// *ast.AssignStmt {
			// Lhs: []ast.Expr (len = 1) {
			// .  0: *ast.Ident {
			// .  .  Name: "x"
			// .  }
			// }
			// Tok: =
			// Rhs: []ast.Expr (len = 1) {
			// .  0: *ast.BasicLit {
			// .  .  Kind: INT
			// .  .  Value: "0"
			// .  }
			// }

			// Assign (
			//   targets =  [
			//     Tuple (
			//       elts =  [
			//         Name (
			//         ) // Name
			//         Name (
			//         ) // Name
			//       ] //
			//     ) // Tuple
			//   ] //
			//   value = ...
			// )

			var left, right []goast.Expr
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "targets"); ok {
						var errLeft error
						left, errLeft = transpileExprs(a.Right.(*List).Args[0])
						if errLeft != nil {
							et.Add(errLeft)
							break
						}
					}
					if _, ok := isIdent(a.Left, "value"); ok {
						var errRight error
						right, errRight = transpileExprs(a.Right)
						if errRight != nil {
							et.Add(errRight)
							break
						}
					}
				}
			}
			if !et.IsError() {
				stmt = &goast.AssignStmt{
					Lhs: left,
					Tok: token.ASSIGN,
					Rhs: right,
				}
			}

		case "Return":
			ret := &goast.ReturnStmt{}
			expr, errExpr := transpileExprs(v.Args[0])
			if errExpr != nil {
				et.Add(errExpr)
				break
			}
			ret.Results = append(ret.Results, expr...)
			stmt = ret

		case "Expr":
			//	Expr (
			//		value = Call (...)
			//	)
			expr, errExpr := transpileExprs(v.Args[0])
			if errExpr != nil {
				et.Add(errExpr)
				break
			}
			if len(expr) == 1 {
				stmt = &goast.ExprStmt{X: expr[0]}
			}

		case "For":
			// Python: for iData in sData:
			//
			// AST:
			// For (
			//   target = Name (
			//     id = 'iData'
			//   )
			//   iter = Name (
			//     id = 'sData'
			//   )
			//   body =  [
			//   ]
			// )
			//
			// type RangeStmt struct {
			//     Key, Value Expr
			//     Tok        token.Token
			//     X          Expr
			//     Body       *BlockStmt
			// }
			rangestmt := goast.RangeStmt{
				Tok: token.DEFINE,
			}
			for i := range v.Args {
				if a, ok := isAssign(v.Args[i], "target"); ok {
					expr, errExpr := transpileExprs(a.Right)
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					if len(expr) != 1 {
						continue
					}
					rangestmt.Key = expr[0]
				}
				if a, ok := isAssign(v.Args[i], "iter"); ok {
					expr, errExpr := transpileExprs(a.Right)
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					if len(expr) != 1 {
						continue
					}
					rangestmt.X = expr[0]
				}
				if a, ok := isAssign(v.Args[i], "body"); ok {
					stmts, errStmts := transpileStmts(a.Right)
					if errStmts != nil {
						et.Add(errStmts)
						break
					}
					rangestmt.Body = &goast.BlockStmt{
						List: stmts,
					}
				}
			}
			if !et.IsError() {
				stmt = &rangestmt
			}

		case "While":
			// While (
			//   test = Compare (...)
			//   body =  [ ... ]
			// ) // While
			//
			// *ast.ForStmt {
			// Cond: ast.Expr{},
			// Body: *ast.BlockStmt{}
			// }
			var forstmt goast.ForStmt
			for i := range v.Args {
				if a, ok := isAssign(v.Args[i], "test"); ok {
					expr, errExpr := transpileExprs(a.Right)
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					if len(expr) != 1 {
						continue
					}
					forstmt.Cond = expr[0]
				}
				if a, ok := isAssign(v.Args[i], "body"); ok {
					stmts, errStmts := transpileStmts(a.Right)
					if errStmts != nil {
						et.Add(errStmts)
						break
					}
					forstmt.Body = &goast.BlockStmt{
						List: stmts,
					}
				}
			}
			if !et.IsError() {
				stmt = &forstmt
			}

		case "If":
			// From:
			//
			// If (
			//   test = Name (
			//     id = 'False'
			//   )
			//   body   =  [ ... ]
			//   orelse =  [ ... ]
			// )
			//
			// To:
			//
			// *ast.IfStmt {
			//    Cond: *ast.BinaryExpr {
			//       X: {}
			//       Op: ==
			//       Y: {}
			//    }
			//    Body: *ast.BlockStmt { }
			//    Else: *ast.BlockStmt { }
			// }
			ifs := &goast.IfStmt{}
			for i := range v.Args {
				a, ok := v.Args[i].(*Assign)
				if !ok {
					continue
				}
				id, ok := a.Left.(*Ident)
				if !ok {
					continue
				}
				switch id.Name {
				case "test":
					expr, errExpr := transpileExprs(a.Right)
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					if len(expr) != 1 {
						continue
					}
					ifs.Cond = expr[0]
				case "body":
					stmts, errStmts := transpileStmts(a.Right)
					if errStmts != nil {
						et.Add(errStmts)
						break
					}
					ifs.Body = &goast.BlockStmt{
						List: stmts,
					}
				case "orelse":
					stmts, errStmts := transpileStmts(a.Right)
					if errStmts != nil {
						et.Add(errStmts)
						break
					}
					if 0 < len(stmts) {
						ifs.Else = &goast.BlockStmt{
							List: stmts,
						}
					}
				}
			}
			if !et.IsError() {
				stmt = ifs
			}

		default:
			call := &goast.CallExpr{}
			call.Fun = goast.NewIdent(v.Name)
			isCall := false
			for i := range v.Args {
				if a, ok := isAssign(v.Args[i], "values"); ok {
					expr, errExpr := transpileExprs(a.Right)
					isCall = true
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					call.Args = append(call.Args, expr...)
				}
			}
			if !et.IsError() && isCall {
				stmt = &goast.ExprStmt{X: call}
			}
		}

	default:
		et.Add(fmt.Errorf("cannot transpile : %s", n))
	}
	if et.IsError() {
		err = et
		return
	}
	return
}

// func transpileExprs(n Node) (expr []goast.Expr, err error) {
// 	et := errors.New("Error in func transpileExpr")
// 	l, ok := n.(*List)
// 	if !ok {
// 		return
// 	}
// 	for i := range l.Args {
// 		ll, ok := l.Args[i].(*List)
// 		if !ok {
// 			continue
// 		}
// 		//   [
// 		//     Name ( ... )
// 		//   ]
// 		n, en := transpileExprs(ll.Args[0])
// 		if en != nil {
// 			et.Add(en)
// 			continue
// 		}
// 		if !et.IsError() {
// 			expr = append(expr, n...)
// 		}
// 	}
//
// 	if et.IsError() {
// 		err = et
// 		return
// 	}
// 	return
//
// }

func transpileExprs(n Node) (exprs []goast.Expr, err error) {
	et := errors.New("Error in func transpileExprs")

	trs := []func() ([]goast.Expr, bool, error){

		func() (exprs []goast.Expr, ok bool, err error) {
			a, oka := isAssign(n, "value")
			if !oka {
				return nil, false, nil
			}
			ok = true
			exprs, err = transpileExprs(a.Right)
			if err != nil {
				et.Add(err)
			} else {
				s, errs := exprsSel(exprs)
				if errs != nil {
					err = errs
				} else {
					exprs = append([]goast.Expr{}, s)
				}
			}
			return
		},

		func() (exprs []goast.Expr, ok bool, err error) {
			a, oka := isAssign(n, "attr")
			if !oka {
				return nil, false, nil
			}
			ok = true
			name := clearName(fmt.Sprintf("%s", a.Right))
			exprs = append(exprs, goast.NewIdent(name))
			return
		},

		func() (exprs []goast.Expr, ok bool, err error) {
			a, oka := isAssign(n, "name")
			if !oka {
				return nil, false, nil
			}
			ok = true
			name := clearName(fmt.Sprintf("%s", a.Right))
			exprs = append(exprs, goast.NewIdent(name))
			return
		},

		func() (exprs []goast.Expr, ok bool, err error) {
			list, okl := isList(n, "Tuple")
			if !okl {
				return nil, false, nil
			}
			// Tuple (
			//   elts =  [
			//     Name ( ... )
			//     Name ( ... )
			//   ] //
			// )
			ok = true
			for i := range list.Args {
				a, ok := isAssign(list.Args[i], "elts")
				if !ok {
					// ignored elements
					continue
				}
				listin := a.Right.(*List)
				for j := range listin.Args {
					ep, errp := transpileExprs(listin.Args[j])
					if errp != nil {
						return nil, true, errp
					}
					exprs = append(exprs, ep...)
				}
			}
			return
		},

		func() (exprs []goast.Expr, ok bool, err error) {
			list, okl := isList(n, "Name")
			if !okl {
				return nil, false, nil
			}
			// Name (
			//   id = 'BaseModule'
			//   ctx = Load ()
			// )
			ok = true
			a, oka := isAssign(list.Args[0], "id")
			if !oka {
				err = fmt.Errorf("Strange: %s", n)
				return
			}
			name := clearName(fmt.Sprintf("%s", a.Right))
			exprs = append(exprs, goast.NewIdent(name))
			return
		},

		func() (exprs []goast.Expr, ok bool, err error) {
			v, ok := n.(*List)
			if !ok {
				return
			}

			if len(v.Args) == 1 {
				if a, ok := v.Args[0].(*Assign); ok {
					if _, ok := isIdent(a.Left, "n"); ok {
						exprs = append(exprs, goast.NewIdent(a.Right.(*Ident).Name))
						return exprs, true, nil
					}
				}
			}

			defer func() {
				if err != nil {
					err = fmt.Errorf("Name: %v\n%v", v.Name, err)
				}
			}()

			switch v.Name {
			case "":

				for i := range v.Args {
					es, ers := transpileExprs(v.Args[i])
					if ers != nil {
						et.Add(ers)
						continue
					}
					exprs = append(exprs, es...)
				}

			case "Expr":
				//	Expr (
				//		value = Call (...)
				//	)
				exprs, err = transpileExprs(v.Args[0])
				return exprs, true, err

			case "Subscript":
				// Subscript (
				//   value = Name ( ... )
				//   slice = Slice (
				//      lower = None
				//      upper = None
				//      step = None
				//   ) // Slice
				// ) // Subscript
				var e goast.Expr
				var x goast.Expr
				for i := range v.Args {
					if a, ok := isAssign(v.Args[i], "value"); ok {
						ea, erra := transpileExprs(a.Right)
						if erra != nil {
							et.Add(erra)
							continue
						}
						x, erra = exprsSel(ea)
						if erra != nil {
							et.Add(erra)
							continue
						}
					}
					if a, ok := isAssign(v.Args[i], "slice"); ok {
						if _, ok := isList(a.Right, "Slice"); ok {
							// Slice (
							//   lower = None
							//   upper = None
							//   step = None
							// ) // Slice
							e = &goast.SliceExpr{
								X: x,
							}
						} else {
							// Index (
							//   value = Call (
							//   ) // Call
							// ) // Index
							ea, erra := transpileExprs(a.Right.(*List).Args[0])
							if erra != nil {
								et.Add(erra)
								continue
							}
							indexs, erra := exprsSel(ea)
							if erra != nil {
								et.Add(erra)
								continue
							}
							e = &goast.IndexExpr{
								X:     x,
								Index: indexs,
							}
						}
					}
				}
				if !et.IsError() {
					exprs = append(exprs, e)
				}

				// UnaryOp (
				//    op = Not ( )
				//    operand = ...
				// )
			case "UnaryOp":
				unary := goast.UnaryExpr{}
				for i := range v.Args {
					a, ok := v.Args[i].(*Assign)
					if !ok {
						continue
					}
					_, ok1 := isIdent(a.Left, "op")
					_, ok2 := isIdent(a.Left, "operand")
					if !(ok1 || ok2) {
						continue
					}
					if ok1 {
						tok, errp := transpileOp(a.Right)
						if errp != nil {
							et.Add(errp)
						}
						if !et.IsError() {
							unary.Op = tok
						}
					}
					if ok2 {
						exprp, errp := transpileExprs(a.Right)
						if errp != nil {
							et.Add(errp)
						}
						if !et.IsError() {
							unary.X = exprp[0]
						}
					}
				}
				if !et.IsError() {
					exprs = append(exprs, &unary)
				}

				// Compare (
				//   left = ...
				//   ops  = ...
				//   comparators = ...
				// ) // Compare
			case "Compare":
				bin := goast.BinaryExpr{}
				for i := range v.Args {
					a, ok := v.Args[i].(*Assign)
					if !ok {
						continue
					}
					if _, ok := isIdent(a.Left, "left"); ok {
						ep, errp := transpileExprs(a.Right)
						if errp != nil {
							et.Add(errp)
							break
						}
						bin.X = ep[0]
					}
					if _, ok := isIdent(a.Left, "ops"); ok {
						tok, errp := transpileOp(a.Right.(*List).Args[0])
						if errp != nil {
							et.Add(fmt.Errorf("Error in Operation: %s\n%v",
								v,
								errp))
							break
						}
						bin.Op = tok
					}
					if _, ok := isIdent(a.Left, "comparators"); ok {
						ep, errp := transpileExprs(a.Right.(*List).Args[0])
						if errp != nil {
							et.Add(errp)
							break
						}
						if len(ep) == 1 {
							bin.Y = ep[0]
						}
					}
				}
				if bin.X == nil || bin.Y == nil {
					et.Add(fmt.Errorf("Nil binary: %s", n))
				}
				if !et.IsError() {
					exprs = append(exprs, &bin)
				}

				// Str (
				//   s = 'text'
				// ) // Str
			case "Str":
				exprs = append(exprs,
					goast.NewIdent("\""+clearName(v.Args[0].(*Assign).Right.(*Ident).Name)+"\""))

			case "Call":
				// Call (
				//   func = Name (
				//     id = 'len'
				//   )
				//   args =  [ ... ]
				// ) // Call
				call := goast.CallExpr{}
				call.Fun = goast.NewIdent(v.Name)
				for i := range v.Args {
					if a, ok := isAssign(v.Args[i], "func"); ok {
						expr, errExpr := transpileExprs(a.Right)
						if errExpr != nil {
							et.Add(errExpr)
							break
						}
						call.Fun, errExpr = exprsSel(expr)
						if errExpr != nil {
							et.Add(errExpr)
							break
						}
					}
					if a, ok := isAssign(v.Args[i], "args"); ok {
						expr, errExpr := transpileExprs(a.Right)
						if errExpr != nil {
							et.Add(errExpr)
							break
						}
						call.Args = append(call.Args, expr...)
					}
				}
				if !et.IsError() {
					exprs = append(exprs, &call)
				}

				//  Attribute (
				//    value = Name (
				//    	id = 'self'
				//    ) // Name
				//    attr = 'tol'
				//  ) // Attribute
			case "Attribute":
				name := ""
				for i := range v.Args {
					a, ok := v.Args[i].(*Assign)
					if !ok {
						continue
					}
					_, ok1 := isIdent(a.Left, "value") // TODO : Subscript!!!
					_, ok2 := isIdent(a.Left, "attr")
					if !(ok1 || ok2) {
						continue
					}
					ea, erra := transpileExprs(v.Args[i])
					if erra != nil {
						et.Add(erra)
						continue
					}
					for i := range ea {
						var buf bytes.Buffer
						printer.Fprint(&buf, token.NewFileSet(), ea[i])
						name += buf.String() + "."
					}
				}
				if !et.IsError() {
					if name[len(name)-1] == '.' {
						// for avoid last point
						name = name[:len(name)-1]
					}
					exprs = append(exprs, goast.NewIdent(name))
				}

			case "BoolOp":
				// BoolOp (
				//   op = Or()
				//   values =  [
				//      Compare(...)
				//      Compare(...)
				//      Compare(...)
				//   ]
				// )
				var (
					cmps []goast.Expr
					op   token.Token
				)
				for i := range v.Args {
					if a, ok := isAssign(v.Args[i], "op"); ok {
						opp, errp := transpileOp(a.Right)
						if errp != nil {
							et.Add(errp)
							continue
						}
						op = opp
					}
					if a, ok := isAssign(v.Args[i], "values"); ok {
						exprs, errExpr := transpileExprs(a.Right)
						if errExpr != nil {
							et.Add(errExpr)
							continue
						}
						cmps = exprs
					}
				}
				if len(cmps) < 2 {
					et.Add(fmt.Errorf("Compare is not enought"))
				} else {
					bin := &goast.BinaryExpr{}
					bin.X = cmps[0]
					bin.Op = op
					bin.Y = cmps[1]
					for i := 2; i < len(cmps); i++ {
						bin = &goast.BinaryExpr{
							X:  bin,
							Op: op,
							Y:  cmps[i],
						}
					}
					exprs = append(exprs, bin)
				}

			case "BinOp":
				// BinOp (
				//   left = ...
				//   op = ...
				//   right = ...
				// ) // BinOp
				bin := goast.BinaryExpr{}
				for i := range v.Args {
					a, ok := v.Args[i].(*Assign)
					if !ok {
						continue
					}
					if _, ok := isIdent(a.Left, "left"); ok {
						ep, errp := transpileExprs(a.Right)
						if errp != nil {
							et.Add(errp)
							break
						}
						bin.X = ep[0]
					}
					if _, ok := isIdent(a.Left, "op"); ok {
						tok, errp := transpileOp(a.Right)
						if errp != nil {
							et.Add(fmt.Errorf("Error in Operation: %s\n%v",
								v,
								errp))
							break
						}
						bin.Op = tok
					}
					if _, ok := isIdent(a.Left, "right"); ok {
						ep, errp := transpileExprs(a.Right)
						if errp != nil {
							et.Add(errp)
							break
						}
						bin.Y, errp = exprsSel(ep)
						if errp != nil {
							et.Add(errp)
							break
						}
					}
				}

				if bin.X == nil || bin.Y == nil {
					et.Add(fmt.Errorf("Not valid BinaryOp: %s", n))
					break
				}

				// Power exception
				if bin.Op == POW {
					// *ast.CallExpr {
					// Fun: *ast.SelectorExpr {
					// .  X: *ast.Ident {
					// .  .  Name: "math"
					// .  }
					// .  Sel: *ast.Ident {
					// .  .  Name: "Pow"
					// .  }
					// }
					// Args: []ast.Expr (len = 2) {
					// .  0: *ast.Ident {
					// .  .  Name: "x"
					// .  }
					// .  1: *ast.Ident {
					// .  .  Name: "y"
					// .  }
					// }
					exprs = append(exprs, &goast.CallExpr{
						Fun:  goast.NewIdent("math.Pow"),
						Args: []goast.Expr{bin.X, bin.Y},
					})
				} else {
					if !et.IsError() {
						exprs = append(exprs, &bin)
					}
				}

			default:
				err = fmt.Errorf("cannot transpile : %s", n)
			}
			return
		},
	}

	for i := range trs {
		es, ok, erre := trs[i]()
		if !ok {
			continue
		}
		if erre != nil {
			et.Add(fmt.Errorf("Error function: %d", i))
			et.Add(erre)
		} else {
			exprs = append(exprs, es...)
		}
		break
	}

	if et.IsError() {
		err = et
		return
	}
	return
}

const (
	// python specific token
	POW token.Token = 100000
)

func transpileOp(n Node) (tok token.Token, err error) {
	et := errors.New("Error in func transpileOp")
	l, ok := n.(*List)
	if !ok {
		err = fmt.Errorf("not valid List: %s", n)
		return
	}
	switch l.Name {
	case "Not":
		tok = token.NOT
	case "Mult":
		tok = token.MUL
	case "Add":
		tok = token.ADD
	case "Pow":
		tok = POW
	case "Eq", "Is":
		tok = token.EQL
	case "NotEq":
		tok = token.NEQ
	case "Div":
		tok = token.QUO
	case "Or":
		tok = token.LOR
	case "Gt":
		tok = token.GTR
	case "Lt":
		tok = token.LSS
	case "Sub", "USub":
		tok = token.SUB
	case "Mod":
		tok = token.ADD // concat for strings
	default:
		tok = token.AND_NOT_ASSIGN // some trash token
		et.Add(fmt.Errorf("not valid token: %s for %s", l.Name, n))
	}
	if et.IsError() {
		err = et
		return
	}
	return
}
