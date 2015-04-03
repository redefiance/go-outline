package outline

import (
	"fmt"
	"strings"
)

type Decl struct {
	Token string
	Vars  []Variable

	LineFrom, LineTo int
}

func (decl Decl) String() string {
	var s []string
	for _, v := range decl.Vars {
		s = append(s, v.Name)
	}
	return fmt.Sprintf("%s %s", decl.Token, strings.Join(s, ", "))
}
