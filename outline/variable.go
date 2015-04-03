package outline

import (
	"fmt"
	"strings"
)

type Variables []Variable

func (vs Variables) String() string {
	var s []string
	for _, v := range vs {
		s = append(s, v.String())
	}
	return strings.Join(s, ", ")
}

type Variable struct {
	Name string
	Type Type
}

func (v Variable) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Type)
}
