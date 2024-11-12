package cmd

import (
	"fanes-cli/cmd/generate"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fanes",
	Short: "CLI for Fanes",
	Long:  `CLI that helps Fanes with basic utilities`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello ( •̀ ω •́ )✧, type fanes --help for more information!")
	},
}

func Execute() {
	generate.Init(rootCmd)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
