package outline

import (
	"fmt"
	"strings"
)

func list(things ...fmt.Stringer) string {
	var strs []string
	for _, thing := range things {
		strs = append(strs, thing.String())
	}
	return strings.Join(strs, ",")
}

//
// import (
// 	"fmt"
// 	"go/ast"
// 	"reflect"
// )
//
// func (p *Package) Visit(node ast.Node) ast.Visitor {
// 	switch e := node.(type) {
//
// 	case nil, *ast.Ident, *ast.ImportSpec:
// 		return nil
//
// 	case *ast.Package, *ast.File, *ast.GenDecl:
// 		return p
//
// 	case *ast.TypeSpec:
// 		return nil
//
// 	case *ast.ValueSpec:
// 		t := parseType(e)
// 		for _, ident := range e.Names {
// 			p.Vars = append(p.Vars, &Variable{
// 				Name: ident.Name,
// 				Type: t,
// 			})
// 		}
// 		return nil
//
// 	case *ast.FuncDecl:
// 		return nil
//
// 	default:
// 		fmt.Println("package: unhandled node", reflect.TypeOf(e))
// 		return nil
// 	}
// }
//
// func (p *Package) parseType(node ast.Node) Type {
// 	switch e := node.(type) {
//
// 	case *ast.Ident:
// 		return p.getType(e.Name)
//
// 	case *ast.ArrayType:
// 		n := uint(0)
// 		if e.Len != nil {
// 			fmt.Println(e)
// 		}
// 		return ArrayType{
// 			Len:  n,
// 			Type: parseType(e.Elt),
// 		}
//
// 	case *ast.CallExpr:
// 		fmt.Println(e)
// 		return nil
//
// 	case *ast.CompositeLit:
// 		return parseType(e.Type)
//
// 	case *ast.ValueSpec:
// 		if e.Type != nil {
// 			return parseType(e.Type)
// 		}
// 		return parseType(e.Values[0])
//
// 	default:
// 		fmt.Println("type: unhandled node", reflect.TypeOf(e))
// 		return nil
// 	}
// }
//
// // func (p Type) Visit(node ast.Node) ast.Visitor {
// // 	return nil
// // }
//
// type valueParser struct{}
//
// func (p valueParser) Visit(node ast.Node) ast.Visitor {
// 	fmt.Println("<<NEW VALUE>>")
//
// 	switch e := node.(type) {
//
// 	case *ast.Ident:
// 		// v := Variable{Name: e.Name}
// 		return nil
//
// 	case *ast.CompositeLit:
// 		fmt.Println(e)
// 		return nil
//
// 	default:
// 		fmt.Println("unhandled value type", reflect.TypeOf(e))
// 		return nil
// 	}
// }
//
// type funcParser struct{}
//
// func (p funcParser) Visit(node ast.Node) ast.Visitor {
// 	return nil
// }
