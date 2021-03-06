package ast

import (
	"testing"
)

func TestRecordDecl(t *testing.T) {
	nodes := map[string]Node{
		`0x7f913c0dbb50 <line:76:9, line:79:1> line:76:9 union definition`: &RecordDecl{
			Address:    "0x7f913c0dbb50",
			Position:   "line:76:9, line:79:1",
			Prev:       "",
			Position2:  "line:76:9",
			Kind:       "union",
			Name:       "",
			Definition: true,
			Children:   []Node{},
		},
		`0x7f85360285c8 </usr/include/sys/_pthread/_pthread_types.h:57:1, line:61:1> line:57:8 struct __darwin_pthread_handler_rec definition`: &RecordDecl{
			Address:    "0x7f85360285c8",
			Position:   "/usr/include/sys/_pthread/_pthread_types.h:57:1, line:61:1",
			Prev:       "",
			Position2:  "line:57:8",
			Kind:       "struct",
			Name:       "__darwin_pthread_handler_rec",
			Definition: true,
			Children:   []Node{},
		},
		`0x7f85370248a0 <line:94:1, col:8> col:8 struct __sFILEX`: &RecordDecl{
			Address:    "0x7f85370248a0",
			Position:   "line:94:1, col:8",
			Prev:       "",
			Position2:  "col:8",
			Kind:       "struct",
			Name:       "__sFILEX",
			Definition: false,
			Children:   []Node{},
		},
	}

	runNodeTests(t, nodes)
}
