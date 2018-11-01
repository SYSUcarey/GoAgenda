package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deluserCmd represents the deluser command
var deluserCmd = &cobra.Command{
	Use:   "deluser",
	Short: "删除本用户",
	Long: `
GoAgenda deluser -p password
	
	各个参数分别对应:
	-p 用户密码`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		password, _ := cmd.Flags().GetString("pass")
		// 处理参数
		if password == "" {
			fmt.Println("password cannot be null")
			return
		}
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(deluserCmd)


	deluserCmd.Flags().StringP("pass", "p", "", "your password")
}
