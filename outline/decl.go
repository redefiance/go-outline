package outline

import (
	"fmt"
	"strings"
)

// Decl is a top level declaration. Token can be 'const', 'var', 'type' or 'func'.
// If Decl.Variable == nil, it is a block containing one or more declarations of
// the same type.
type Decl struct {
	DeclBlock
	Token string
	Group []DeclBlock
}

// A DeclBlock is the declaration of one more Variables within a block of code
type DeclBlock struct {
	FilePosition
	Vars []Variable
}

// FilePosition describes a range of lines within a source file
type FilePosition struct {
	LineFrom, LineTo int
}

func (d Decl) String() string {
	s := ""
	for _, g := range d.Group {
		s += g.String()
	}
	if len(s) > 0 {
		return fmt.Sprintf("group %d %s%s%s", len(d.Group), d.Token, d.FilePosition, s)
	}
	return fmt.Sprintf("%s %s", d.Token, d.DeclBlock)
}

func (d DeclBlock) String() string {
	var s []string
	for _, v := range d.Vars {
		s = append(s, v.Name)
	}
	return fmt.Sprintf("%s%s", strings.Join(s, ", "), d.FilePosition)
}

func (fp FilePosition) String() string {
	return fmt.Sprintf(":%d:%d\n", fp.LineFrom, fp.LineTo)
}
