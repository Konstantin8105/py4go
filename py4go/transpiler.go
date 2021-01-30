package py4go

import (
	"bytes"
	"fmt"
	goast "go/ast"
	"go/format"
	"go/token"
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
		mainDecl := goast.FuncDecl{
			Name: goast.NewIdent("main"),
			Type: &goast.FuncType{
				Params: &goast.FieldList{},
			},
			Body: &goast.BlockStmt{List: stmts},
		}
		decls = append(decls, &mainDecl)
	}

	gonode.Decls = decls

	// Create a FileSet for node. Since the node does not come
	// from a real source file, fset will be empty.
	fset := token.NewFileSet()

	var buf bytes.Buffer
	defer func() {
		if r := recover(); r != nil {
			_ = goast.Print(fset, gonode)
			err = fmt.Errorf("%#v", r)
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

func transpile(n Node) (decls []goast.Decl, stmts []goast.Stmt, err error) {
	et := errors.New("Error in func transpile")

	addToMainStmt := func(n Node) {
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
		case "FunctionDef":
			decl := goast.FuncDecl{
				Type: &goast.FuncType{
					Params: &goast.FieldList{},
				},
				Body: &goast.BlockStmt{},
			}
			for i := range v.Args {
				switch vv := v.Args[i].(type) {
				case *Assign:
					if _, ok := isIdent(vv.Left, "name"); ok {
						name := fmt.Sprintf("%s", vv.Right)
						decl.Name = goast.NewIdent(clearName(name))
					}
					if _, ok := isIdent(vv.Left, "body"); ok {
						var (
							errBody error
							body    []goast.Stmt
						)
						body, errBody = transpileStmts(vv.Right)
						if errBody != nil {
							et.Add(errBody)
							continue
						}
						decl.Body.List = append(decl.Body.List, body...)
					}

				default:
				}
			}
			decls = append(decls, &decl)
		default:
			addToMainStmt(n)
			// et.Add(fmt.Errorf("cannot transpile List : %s", n))
		}
	default:
		addToMainStmt(n)
		//et.Add(fmt.Errorf("cannot transpile : %s", n))
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
				et.Add(errStmt)
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
	et := errors.New("Error in func transpileStmt")
	switch v := n.(type) {
	case *Ident:
		if v.Name == "" {
			// ignore empty ident
			break
		}
	case *List:
		switch v.Name {
		case "Assign":
			// from:
			//
			// Assign (
			//   targets =  [
			//     Name (
			//       id = 'x'
			//       ctx = Store (
			//
			//       ) // Store
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
			var left, right goast.Expr
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "targets"); ok {
						var errLeft error
						left, errLeft = transpileExpr(a.Right.(*List).Args[0])
						if errLeft != nil {
							et.Add(errLeft)
							break
						}
					}
					if _, ok := isIdent(a.Left, "value"); ok {
						var errRight error
						right, errRight = transpileExpr(a.Right)
						if errRight != nil {
							et.Add(errRight)
							break
						}
					}
				}
			}
			stmt = &goast.AssignStmt{
				Lhs: []goast.Expr{left},
				Tok: token.ASSIGN,
				Rhs: []goast.Expr{right},
			}

		case "Return":
			ret := &goast.ReturnStmt{}
			expr, errExpr := transpileExpr(v.Args[0])
			if errExpr != nil {
				et.Add(errExpr)
				break
			}
			ret.Results = append(ret.Results, expr)
			stmt = ret

		case "Expr":
			//	Expr (
			//		value = Call (...)
			//	)
			expr, errExpr := transpileExpr(v.Args[0])
			if errExpr != nil {
				et.Add(errExpr)
				break
			}
			stmt = &goast.ExprStmt{X: expr}

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
					expr, errExpr := transpileExpr(a.Right)
					if errExpr != nil {
						et.Add(errExpr)
						break
					}
					ifs.Cond = expr
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
			stmt = ifs

		default:
			call := &goast.CallExpr{}
			call.Fun = goast.NewIdent(v.Name)
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "values"); ok {
						switch v := a.Right.(type) {
						case *List:
							for j := range v.Args {
								expr, errExpr := transpileExpr(v.Args[j])
								if errExpr != nil {
									et.Add(errExpr)
									break
								}
								call.Args = append(call.Args, expr)
							}
						default:
							et.Add(fmt.Errorf("%s", n))
						}
					}
				}
			}
			stmt = &goast.ExprStmt{X: call}
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

func transpileExpr(n Node) (expr goast.Expr, err error) {
	et := errors.New("Error in func transpileExpr")
	switch v := n.(type) {
	case *Assign:
		if _, ok := isIdent(v.Left, "value"); ok {
			return transpileExpr(v.Right)
		}
	case *List:
		if len(v.Args) == 1 {
			if a, ok := v.Args[0].(*Assign); ok {
				if _, ok := isIdent(a.Left, "n"); ok {
					expr = goast.NewIdent(a.Right.(*Ident).Name)
					break
				}
			}
		}
		if v.Name == "Name" {
			if a, ok := v.Args[0].(*Assign); ok {
				id := a.Right.(*Ident) // id
				expr = goast.NewIdent(clearName(id.Name))
				break
			}
		}

		//	Expr (
		//		value = Call (...)
		//	)
		if v.Name == "Expr" {
			return transpileExpr(v.Args[0])
		}

		// Call (
		//   func = Name (
		//     id = 'print_x'
		//     ctx = Load (
		//     ) // Load
		//   ) // Name
		//   args =  [
		//   ] //
		//   keywords =  [
		//   ] //
		//   starargs = None
		//   kwargs = None
		// ) // Call
		if v.Name == "Call" {
			callExpr:= goast.CallExpr{
			}
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if _, ok := isIdent(a.Left, "func"); ok {
						var errFun error
						if errFun != nil {
							et.Add(errFun)
							break
						}
						callExpr.Fun, errFun = transpileExpr(a.Right)
					}
				}
			}
			expr = &callExpr
			break
		}

		et.Add(fmt.Errorf("List : %s", n))
	default:
		et.Add(fmt.Errorf("cannot transpile : %s", n))
	}
	if et.IsError() {
		err = et
		return
	}
	return

}

// /home/konstantin/go/src/github.com/Konstantin8105/py4go/py4go/../testdata/p.py
// Module (
//   body =  [
//     FunctionDef (
//       name = 'pi'
//       args = arguments (
//         args =  [
//
//         ] //
//         vararg = None
//         kwarg = None
//         defaults =  [
//
//         ] //
//       ) // arguments
//       body =  [
//         Return (
//           value = Num (
//             n = 3.1415
//           ) // Num
//         ) // Return
//       ] //
//       decorator_list =  [
//
//       ] //
//     ) // FunctionDef
//     Print (
//       dest = None
//       values =  [
//         BinOp (
//           left = Num (
//             n = 2
//           ) // Num
//           op = Mult (
//
//           ) // Mult
//           right = Call (
//             func = Name (
//               id = 'pi'
//               ctx = Load (
//
//               ) // Load
//             ) // Name
//             args =  [
//
//             ] //
//             keywords =  [
//
//             ] //
//             starargs = None
//             kwargs = None
//           ) // Call
//         ) // BinOp
//       ] //
//       nl = True
//     ) // Print
//     Assign (
//       targets =  [
//         Name (
//           id = 'x'
//           ctx = Store (
//
//           ) // Store
//         ) // Name
//       ] //
//       value = Num (
//         n = 1
//       ) // Num
//     ) // Assign
//     FunctionDef (
//       name = 'print_x'
//       args = arguments (
//         args =  [
//
//         ] //
//         vararg = None
//         kwarg = None
//         defaults =  [
//
//         ] //
//       ) // arguments
//       body =  [
//         Print (
//           dest = None
//           values =  [
//             Name (
//               id = 'x'
//               ctx = Load (
//
//               ) // Load
//             ) // Name
//           ] //
//           nl = True
//         ) // Print
//         If (
//           test = Name (
//             id = 'False'
//             ctx = Load (
//
//             ) // Load
//           ) // Name
//           body =  [
//             Assign (
//               targets =  [
//                 Name (
//                   id = 'x'
//                   ctx = Store (
//
//                   ) // Store
//                 ) // Name
//               ] //
//               value = Num (
//                 n = 0
//               ) // Num
//             ) // Assign
//           ] //
//           orelse =  [
//
//           ] //
//         ) // If
//       ] //
//       decorator_list =  [
//
//       ] //
//     ) // FunctionDef
//     Expr (
//       value = Call (
//         func = Name (
//           id = 'print_x'
//           ctx = Load (
//
//           ) // Load
//         ) // Name
//         args =  [
//
//         ] //
//         keywords =  [
//
//         ] //
//         starargs = None
//         kwargs = None
//       ) // Call
//     ) // Expr
//     Assign (
//       targets =  [
//         Name (
//           id = 'x'
//           ctx = Store (
//
//           ) // Store
//         ) // Name
//       ] //
//       value = BinOp (
//         left = BinOp (
//           left = BinOp (
//             left = Name (
//               id = 'x'
//               ctx = Load (
//
//               ) // Load
//             ) // Name
//             op = Add (
//
//             ) // Add
//             right = Num (
//               n = 1
//             ) // Num
//           ) // BinOp
//           op = Mult (
//
//           ) // Mult
//           right = BinOp (
//             left = Name (
//               id = 'x'
//               ctx = Load (
//
//               ) // Load
//             ) // Name
//             op = Add (
//
//             ) // Add
//             right = Num (
//               n = 3
//             ) // Num
//           ) // BinOp
//         ) // BinOp
//         op = Mult (
//
//         ) // Mult
//         right = BinOp (
//           left = Num (
//             n = 5
//           ) // Num
//           op = Pow (
//
//           ) // Pow
//           right = Num (
//             n = 8
//           ) // Num
//         ) // BinOp
//       ) // BinOp
//     ) // Assign
//     Print (
//       dest = None
//       values =  [
//         Name (
//           id = 'x'
//           ctx = Load (
//
//           ) // Load
//         ) // Name
//       ] //
//       nl = True
//     ) // Print
//   ] //
// ) // Module