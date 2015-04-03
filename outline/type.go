package outline

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Type interface {
	fmt.Stringer
}

// ArrayType describes an array of Type. If Len == 0, it is a slice of Type.
type ArrayType struct {
	Len  uint
	Type Type
}

func (t ArrayType) String() string {
	if t.Len == 0 {
		return fmt.Sprintf("[]%s", t.Type)
	}
	return fmt.Sprintf("[%d]%s", t.Len, t.Type)
}

// BuiltInType describes one of the basic types built into Go
type BuiltInType string

func (t BuiltInType) String() string {
	return string(t)
}

// ChannelType describes a channel of Type
type ChannelType struct {
	Type Type
}

func (t ChannelType) String() string {
	return fmt.Sprintf("chan %s", t.Type)
}

// FuncType describes a function
type FuncType struct {
	Inputs  []Variable
	Outputs []Variable
}

func (t FuncType) String() string {
	return fmt.Sprintf("func(%s) %s", "...", "...")
}

type ImportedType struct {
	Package  string
	Typename string
}

func (t ImportedType) String() string {
	return fmt.Sprintf("%s.%s", t.Package, t.Typename)
}

// InterfaceType describes an interface
type InterfaceType struct {
	Methods []Variable
}

func (t InterfaceType) String() string {
	return "interface{ ... }"
}

// MapType describes a map[KeyType]ValueType
type MapType struct {
	KeyType   Type
	ValueType Type
}

func (t MapType) String() string {
	return fmt.Sprintf("map[%s]%s", t.KeyType, t.ValueType)
}

// StructType describes a struct
type StructType struct {
	Fields []Variable
}

func (t StructType) String() string {
	return "struct{ ... }"
}

// StructType describes a pointer
type PointerType struct {
	Type Type
}

func (t PointerType) String() string {
	return fmt.Sprintf("*%s", t.Type)
}

type typePromise string

func (t typePromise) String() string {
	return fmt.Sprintf("unresolved(%s)", string(t))
}

func parseType(node ast.Expr) Type {
	switch e := node.(type) {

	case nil:
		return nil

	case *ast.Ident:
		switch e.Name {
		case "bool":
			return BuiltInType("bool")
		case "byte":
			return BuiltInType("byte")
		case "complex128":
			return BuiltInType("complex128")
		case "complex64":
			return BuiltInType("complex64")
		case "error":
			return BuiltInType("error")
		case "float32":
			return BuiltInType("float32")
		case "float64":
			return BuiltInType("float64")
		case "int":
			return BuiltInType("int")
		case "int16":
			return BuiltInType("int16")
		case "int32":
			return BuiltInType("int32")
		case "int64":
			return BuiltInType("int64")
		case "int8":
			return BuiltInType("int8")
		case "rune":
			return BuiltInType("rune")
		case "string":
			return BuiltInType("string")
		case "uint":
			return BuiltInType("uint")
		case "uint16":
			return BuiltInType("uint16")
		case "uint32":
			return BuiltInType("uint32")
		case "uint64":
			return BuiltInType("uint64")
		case "uint8":
			return BuiltInType("uint8")
		case "uintptr":
			return BuiltInType("uintptr")
		default:
			return typePromise(e.Name)
		}

	case *ast.ChanType:
		return ChannelType{
			Type: parseType(e.Value),
		}

	case *ast.SelectorExpr:
		return ImportedType{
			Package:  e.Sel.Name,
			Typename: e.X.(*ast.Ident).Name,
		}

	case *ast.StarExpr:
		return PointerType{
			Type: parseType(e.X),
		}

	case *ast.ArrayType:
		l := uint(0)
		switch e := e.Len.(type) {
		case nil:
		case *ast.Ident:
		// TODO
		case *ast.SelectorExpr:
		// TODO
		default:
			panic(e)
		}
		return ArrayType{
			Len:  l,
			Type: parseType(e.Elt),
		}

	case *ast.FuncType:
		return FuncType{}

	case *ast.InterfaceType:
		t := InterfaceType{}
		for _, field := range e.Methods.List {
			funcType := parseType(field.Type)
			for _, ident := range field.Names {
				t.Methods = append(t.Methods, Variable{
					Name: ident.Name,
					Type: funcType,
				})
			}
		}
		return t

	case *ast.StructType:
		t := StructType{}
		for _, field := range e.Fields.List {
			fieldtype := parseType(field.Type)
			if field.Names == nil {
				t.Fields = append(t.Fields, Variable{"", fieldtype})
			} else {
				for _, ident := range field.Names {
					t.Fields = append(t.Fields, Variable{ident.Name, fieldtype})
				}
			}
		}
		return t
	}

	fmt.Println("type: unhandled node", reflect.TypeOf(node))
	return nil
}
