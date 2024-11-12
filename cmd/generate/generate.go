package generate

import "github.com/spf13/cobra"

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random things",
	Long:  `The generate command is used to generate random things like id, files, etc.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(GenerateCmd)
}
