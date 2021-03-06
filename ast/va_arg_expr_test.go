package ast

import (
	"testing"
)

func TestVAArgExpr(t *testing.T) {
	nodes := map[string]Node{
		`0x7ff7d314bca8 <col:6, col:31> 'int *'`: &VAArgExpr{
			Address:  "0x7ff7d314bca8",
			Position: "col:6, col:31",
			Type:     "int *",
			Children: []Node{},
		},
	}

	runNodeTests(t, nodes)
}
