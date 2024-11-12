package generate

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	uuidTotal = 1
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate UUID v4",
	Long:  `Generate UUID v4`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < uuidTotal; i++ {
			uuidContent()
		}
	},
}

func uuidContent() {
	fmt.Println(uuid.New())
}

func init() {
	uuidCmd.Flags().IntVarP(&uuidTotal, "total", "t", 1, "Specify the number of UUID v4 to generate")

	GenerateCmd.AddCommand(uuidCmd)
}
