package outline

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type Package struct {
	Name  string
	Path  string
	Scope Scope
	Files []File

	ast *ast.Package
	fs  *token.FileSet
}

func ParsePackage(dirpath string, exportedOnly bool) (Package, error) {
	pkg := Package{}
	pkg.Path = dirpath
	pkg.fs = token.NewFileSet()

	asts, err := parser.ParseDir(pkg.fs, dirpath, nil, 0)
	if err != nil {
		return pkg, err
	}

	for pkgname, pkgast := range asts {
		if strings.HasSuffix(pkgname, "_test") {
			continue
		}

		if pkg.Name != "" {
			return pkg, fmt.Errorf("Found multiple packages %s, %s in %s",
				pkg.Name, pkgname, dirpath)
		}

		pkg.Name = pkgname
		pkg.ast = pkgast
	}

	for filename, f := range pkg.ast.Files {
		pkg.Files = append(pkg.Files, pkg.parseFile(filename, f, exportedOnly))
	}

	return pkg, nil

	// p := &Package{}
	//
	// type decl struct {
	// 	name string
	// 	node ast.Node
	// }
	// var decls []decl
	// for _, f := range f.Files {
	//
	//
	// 	for _, d := range f.Decls {
	// 		switch e := d.(type) {
	// 		case *ast.BadDecl:
	// 		case *ast.GenDecl:
	// 			switch e.Tok {
	// 			case token.IMPORT:
	// 				for _, spec := range e.Specs {
	// 					decls = append(decls, decl{"import", spec})
	// 				}
	// 			case token.CONST:
	// 				for _, spec := range e.Specs {
	// 					decls = append(decls, decl{"const", spec})
	// 				}
	// 			case token.TYPE:
	// 				for _, spec := range e.Specs {
	// 					decls = append(decls, decl{"type", spec})
	// 				}
	// 			case token.VAR:
	// 				for _, spec := range e.Specs {
	// 					decls = append(decls, decl{"var", spec})
	// 				}
	// 			}
	// 		case *ast.FuncDecl:
	// 			decls = append(decls, decl{"func", e})
	// 		default:
	// 			panic(e)
	// 		}
	// 	}
	// }
	//
	// lookup := map[string]Type{}
	// for _, decl := range decls {
	// 	switch decl.name {
	// 	case "type":
	// 		e := decl.node.(*ast.TypeSpec)
	// 		lookup[e.Name.Name] = parseType(e.Type)
	// 	case "func":
	// 		e := decl.node.(*ast.FuncDecl)
	// 		lookup[e.Name.Name] = parseType(e.Type)
	// 	}
	// }
	//
	// var resolve func(Type) Type
	// resolve = func(t Type) Type {
	// 	vlist := func(in []Variable) (out []Variable) {
	// 		for _, v := range in {
	// 			out = append(out, Variable{v.Name, resolve(v.Type)})
	// 		}
	// 		return out
	// 	}
	// 	switch e := (t).(type) {
	// 	case typePromise:
	// 		return lookup[string(e)]
	// 	case ArrayType:
	// 		return ArrayType{e.Len, resolve(e.Type)}
	// 	case BuiltInType:
	// 		return e
	// 	case ChannelType:
	// 		return ChannelType{resolve(e.Type)}
	// 	case FuncType:
	// 		return FuncType{vlist(e.Inputs), vlist(e.Outputs)}
	// 	case InterfaceType:
	// 		return InterfaceType{vlist(e.Methods)}
	// 	case MapType:
	// 		return MapType{resolve(e.KeyType), resolve(e.ValueType)}
	// 	case StructType:
	// 		return StructType{vlist(e.Fields)}
	// 	case PointerType:
	// 		return PointerType{resolve(e.Type)}
	// 	default:
	// 		panic(e)
	// 	}
	// }
	// for n, t := range lookup {
	// 	lookup[n] = resolve(t)
	// }
	//
	// parseValueSpec := func(e *ast.ValueSpec, declname string) {
	// 	typ := parseType(e.Type)
	// 	for _, ident := range e.Names {
	// 		p.Decls = append(p.Decls, Decl{Variable{ident.Name, typ}, declname})
	// 	}
	// }
	//
	// for _, decl := range decls {
	// 	switch decl.name {
	// 	case "import":
	// 	case "type":
	// 		e := decl.node.(*ast.TypeSpec)
	// 		p.Decls = append(p.Decls, Decl{Variable{e.Name.Name, parseType(e.Type)}, decl.name})
	// 	case "func":
	// 		e := decl.node.(*ast.FuncDecl)
	// 		p.Decls = append(p.Decls, Decl{Variable{e.Name.Name, parseType(e.Type)}, decl.name})
	// 	case "const":
	// 		parseValueSpec(decl.node.(*ast.ValueSpec), decl.name)
	// 	case "var":
	// 		parseValueSpec(decl.node.(*ast.ValueSpec), decl.name)
	// 	}
	// }
	//
	// return p
}
