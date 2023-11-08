package cmd

import (
	"fmt"

	"github.com/lucasnhso/sygnus-cli/services"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g"},
	Short:   "Generate modules",

	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}
		if !services.ModuleIsValid(args[0]) {
			return fmt.Errorf("invalid module specified: %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		module := args[0]
		name := args[1]
		switch module {
		case "repository", "usecase", "controller":
			services.GenerateModuleFile(module, name)
		case "resource":
			services.GenerateModuleFile("repository", name)
			services.GenerateModuleFile("usecase", name)
			services.GenerateModuleFile("controller", name)
		}

		fmt.Println("Successfully generated")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
