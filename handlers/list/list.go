package list

import "github.com/fatih/color"

func List() {

	for _, tool := range tools {
		color.Yellow(tool.Name)
	}

}
