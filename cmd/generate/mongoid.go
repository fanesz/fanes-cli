package generate

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	mongoIdTotal = 1
)

var mongoIdCmd = &cobra.Command{
	Use:   "mongoid",
	Short: "Generate MongoID",
	Long:  `Generate MongoID`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < mongoIdTotal; i++ {
			mongoIdContent()
		}
	},
}

func mongoIdContent() {
	timestamp := fmt.Sprintf("%x", time.Now().Unix())
	randomBytes := make([]byte, 12)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(timestamp + hex.EncodeToString(randomBytes)[8:])
}

func init() {
	mongoIdCmd.Flags().IntVarP(&mongoIdTotal, "total", "t", 1, "Specify the number of MongoDB ObjectIDs to generate")

	GenerateCmd.AddCommand(mongoIdCmd)
}
