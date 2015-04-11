package outline

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

// File describes a single go source file
type File struct {
	Path  string
	Decls []Decl
}

func (f File) String() string {
	s := ""
	for _, d := range f.Decls {
		s += d.String()
	}
	return fmt.Sprintf("file %s\n%s", f.Path, s)
}

func (p Package) parseFile(filepath string, f *ast.File, exportedOnly bool) File {
	file := File{Path: filepath}

	exported := func(name string) bool {
		if !exportedOnly {
			return true
		}
		return strings.ToUpper(name[0:1]) == name[0:1]
	}

	block := func(node ast.Node, doc *ast.CommentGroup) DeclBlock {
		b := DeclBlock{}
		b.LineTo = p.fs.Position(node.End()).Line
		if doc != nil {
			node = doc
		}
		b.LineFrom = p.fs.Position(node.Pos()).Line
		return b
	}

	parseDecl := func(e *ast.GenDecl) {
		decl := Decl{DeclBlock: block(e, e.Doc)}
		stuff := func(spec ast.Spec, doc *ast.CommentGroup, names []*ast.Ident) {
			sub := block(spec, doc)
			for _, ident := range names {
				if exported(ident.Name) {
					sub.Vars = append(sub.Vars, Variable{ident.Name, nil})
				}
			}
			if len(sub.Vars) > 0 {
				decl.Group = append(decl.Group, sub)
			}
		}
		switch e.Tok {
		case token.IMPORT:
		case token.CONST:
			decl.Token = "const"
			for _, spec := range e.Specs {
				spec := spec.(*ast.ValueSpec)
				stuff(spec, spec.Doc, spec.Names)
			}
		case token.TYPE:
			decl.Token = "type"
			for _, spec := range e.Specs {
				spec := spec.(*ast.TypeSpec)
				stuff(spec, spec.Doc, []*ast.Ident{spec.Name})
			}
		case token.VAR:
			decl.Token = "var"
			for _, spec := range e.Specs {
				spec := spec.(*ast.ValueSpec)
				stuff(spec, spec.Doc, spec.Names)
			}
		}
		if len(decl.Group) == 0 {
			return
		}
		if !e.Lparen.IsValid() {
			decl.Vars = decl.Group[0].Vars
			decl.Group = nil
		}
		file.Decls = append(file.Decls, decl)
	}

	parseFunc := func(e *ast.FuncDecl) {
		name := e.Name.Name
		if !exported(name) {
			return
		}
		if e.Recv != nil {
			rname := ""
			switch id := e.Recv.List[0].Type.(type) {
			case *ast.StarExpr:
				rname = id.X.(*ast.Ident).Name
			case *ast.Ident:
				rname = id.Name
			}
			if !exported(rname) {
				return
			}
			name = fmt.Sprintf("%s.%s", rname, name)
		}
		decl := Decl{
			DeclBlock: block(e, e.Doc),
			Token:     "func",
		}
		decl.Vars = append(decl.Vars, Variable{name, nil})
		file.Decls = append(file.Decls, decl)
	}

	for _, d := range f.Decls {
		decl := Decl{}
		switch e := d.(type) {
		case *ast.BadDecl:
		case *ast.GenDecl:
			parseDecl(e)
		case *ast.FuncDecl:
			parseFunc(e)
		default:
			panic(e)
		}

		if len(decl.Vars) > 0 {
			file.Decls = append(file.Decls, decl)
		}
	}

	return file
}
