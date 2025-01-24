package lib

func GetTool(name string) (func(), bool) {
	for _, tool := range tools {
		if tool.Name == name {
			return tool.Install, true
		}
	}
	return nil, false
}
