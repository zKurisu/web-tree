package utils

func GetNodeLinks(node Node) []string {
	return node.Link
}

func GetNodeAlias(node Node) []string {
	return node.Alias
}

func GetNodeDesc(node Node) []string {
	return node.Desc
}

func GetNodeIcon(node Node) string {
	return node.Icon
}

func GetNodeLabels(node Node) []string {
	return node.Label
}

func GetNodeStyle(node Node) interface{} {
	return node.Style
}
