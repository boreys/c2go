package ast

type AsmLabelAttr struct {
	Address      string
	Position     string
	FunctionName string
	Children     []Node
}

func parseAsmLabelAttr(line string) *AsmLabelAttr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> \"(?P<function>.+)\"",
		line,
	)

	return &AsmLabelAttr{
		Address:      groups["address"],
		Position:     groups["position"],
		FunctionName: groups["function"],
		Children:     []Node{},
	}
}

func (n *AsmLabelAttr) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
