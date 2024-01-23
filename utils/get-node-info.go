package utils

func (node Node) GetNodeLinks() []string {
	return node.Link
}

func (node Node) GetNodeAlias() []string {
	return node.Alias
}

func (node Node) GetNodeDesc() []string {
	return node.Desc
}

func (node Node) GetNodeIcon() string {
	return node.Icon
}

func (node Node) GetNodeLabels() []string {
	return node.Label
}

func (node Node) GetNodeStyle() interface{} {
	return node.Style
}
