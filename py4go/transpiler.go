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
	decls, errDecls := transpile(nodes)
	defer func() {
		if errDecls != nil {
			err = fmt.Errorf("%v\n%v", err, errDecls)
		}
	}()
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

func transpile(n Node) (decls []goast.Decl, err error) {
	et := errors.New("Error in func transpile")
	switch v := n.(type) {
	case *List:
		switch v.Name {
		case "Module":
			for i := range v.Args {
				if a, ok := v.Args[i].(*Assign); ok {
					if left, ok := a.Left.(*Ident); ok && left.Name == "body" {
						if right, ok := a.Right.(*List); ok {
							for j := range right.Args {
								ds, err := transpile(right.Args[j])
								decls = append(decls, ds...)
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
					if id, ok := vv.Left.(*Ident); ok && id.Name == "name" {
						decl.Name = goast.NewIdent(
							strings.Replace(vv.Right.(*Ident).Name, "'", "", -1))
					}
					if id, ok := vv.Left.(*Ident); ok && id.Name == "body" {
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
			et.Add(fmt.Errorf("cannot transpile List : %s", n))
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
	case *List:
		switch v.Name {
		case "Return":
			ret := &goast.ReturnStmt{}
			expr, errExpr := transpileExpr(v.Args[0])
			if errExpr != nil {
				et.Add(errExpr)
				break
			}
			ret.Results = append(ret.Results, expr)
			stmt = ret

		default:
			et.Add(fmt.Errorf("unknown name : %s %s", v.Name, n))
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
		if id, ok := v.Left.(*Ident); ok {
			switch id.Name {
			case "value":
				return transpileExpr(v.Right)
			}
		}
	case *List:
		if len(v.Args) == 1 {
			if a, ok := v.Args[0].(*Assign); ok {
				if id, ok := a.Left.(*Ident); ok && id.Name == "n" {
					expr = goast.NewIdent(a.Right.(*Ident).Name)
				}
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
