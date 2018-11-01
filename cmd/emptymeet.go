package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// emptymeetCmd represents the emptymeet command
var emptymeetCmd = &cobra.Command{
	Use:   "emptymeet",
	Short: "清空会议",
	Long: `
GoAgenda emptymeet -p password

	各个参数分别对应:
	-p 用户密码`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		password, _ := cmd.Flags().GetString("password")
		// 处理参数
		if password == "" {
			fmt.Println("password cannot be null")
			return
		}
		fmt.Println("password: " + password)		
	},
}

func init() {
	rootCmd.AddCommand(emptymeetCmd)

	emptymeetCmd.Flags().StringP("password", "p", "", "your password")
}
