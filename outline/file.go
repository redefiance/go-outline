package outline

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type File struct {
	Path  string
	Decls []Decl
}

func (p Package) parseFile(filepath string, file *ast.File, exportedOnly bool) File {
	f := File{Path: filepath}

	exported := func(name string) bool {
		return strings.ToUpper(name[0:1]) == name[0:1]
	}

	for _, d := range file.Decls {
		decl := Decl{}
		switch e := d.(type) {
		case *ast.BadDecl:
		case *ast.GenDecl:
			if e.Doc != nil {
				decl.LineFrom = p.fs.Position(e.Doc.Pos()).Line
			} else {
				decl.LineFrom = p.fs.Position(e.Pos()).Line
			}
			decl.LineTo = p.fs.Position(e.End()).Line

			switch e.Tok {
			case token.IMPORT:
				// decl.Token = "import"
				// for _, spec := range e.Specs {
				// 	spec := spec.(*ast.ImportSpec)
				// 	decl.Variables = append(decl.Variables, Variable{
				// 		Name: spec.Name.Name,
				// 		Type: nil,
				// 	})
				// }
			case token.CONST:
				decl.Token = "const"
				for _, spec := range e.Specs {
					spec := spec.(*ast.ValueSpec)
					for _, ident := range spec.Names {
						if !exportedOnly || exported(ident.Name) {
							decl.Vars = append(decl.Vars, Variable{
								Name: ident.Name,
								Type: nil,
							})
						}
					}
				}
			case token.TYPE:
				decl.Token = "type"
				for _, spec := range e.Specs {
					spec := spec.(*ast.TypeSpec)
					if !exportedOnly || exported(spec.Name.Name) {
						decl.Vars = append(decl.Vars, Variable{
							Name: spec.Name.Name,
							Type: nil,
						})
					}
				}
			case token.VAR:
				decl.Token = "var"
				for _, spec := range e.Specs {
					spec := spec.(*ast.ValueSpec)
					for _, ident := range spec.Names {
						if !exportedOnly || exported(ident.Name) {
							decl.Vars = append(decl.Vars, Variable{
								Name: ident.Name,
								Type: nil,
							})
						}
					}
				}
			}
		case *ast.FuncDecl:
			decl.Token = "func"
			if !exportedOnly || exported(e.Name.Name) {
				if e.Doc != nil {
					decl.LineFrom = p.fs.Position(e.Doc.Pos()).Line
				} else {
					decl.LineFrom = p.fs.Position(e.Pos()).Line
				}
				decl.LineTo = p.fs.Position(e.End()).Line

				n := e.Name.Name
				if e.Recv != nil {
					rcv := ""
					switch id := e.Recv.List[0].Type.(type) {
					case *ast.StarExpr:
						rcv = id.X.(*ast.Ident).Name
					case *ast.Ident:
						rcv = id.Name
					}
					if !exported(rcv) {
						continue
					}
					n = fmt.Sprintf("(%s).%s", rcv, n)
				}
				decl.Vars = append(decl.Vars, Variable{
					Name: n,
					Type: nil,
				})
			}
		default:
			panic(e)
		}

		if len(decl.Vars) > 0 {
			f.Decls = append(f.Decls, decl)
		}
	}

	return f
}
