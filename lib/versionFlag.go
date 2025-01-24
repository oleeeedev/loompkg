package lib

func VersionFlag(toolName string) string {
	for _, tool := range tools {
		if tool.Name == toolName {
			return tool.Version
		}
	}
	return ""
}
