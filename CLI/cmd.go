package cli

import (
	"fmt"
	"loompkg/ai"
	"loompkg/handlers"
	"loompkg/handlers/list"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "Loom",
		Short: "A modern package manager",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				handlers.Install(args[0])
			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				handlers.Uninstall(args[0])
			},
		},
		&cobra.Command{
			Use:   "list",
			Short: "List all available packages",
			Run: func(cmd *cobra.Command, args []string) {
				list.List()
			},
		},
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of Loom",
			Run: func(cmd *cobra.Command, args []string) {
				handlers.ShowVersion()
			},
		},
		&cobra.Command{
			Use:   "ai [question]",
			Short: "AI assistant",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				result := ai.ResponseAi(args[0])
				fmt.Println(result)
			},
		},
	)

	return rootCmd
}
