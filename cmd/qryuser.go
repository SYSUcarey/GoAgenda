package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// qryuserCmd represents the qryuser command
var qryuserCmd = &cobra.Command{
	Use:   "qryuser",
	Short: "用户查询",
	Long: `
GoAgenda qryuser
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("qryuser called")
	},
}

func init() {
	rootCmd.AddCommand(qryuserCmd)
}
