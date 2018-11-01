package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// quitmeetCmd represents the quitmeet command
var quitmeetCmd = &cobra.Command{
	Use:   "quitmeet",
	Short: "退出会议",
	Long: `
GoAgenda quitmeet -t title -p password

	各个参数分别对应:
	-t 会议标题
	-p 用户密码`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		password, _ := cmd.Flags().GetString("password")
		//处理参数		
		if title == "" {
			fmt.Println("title cannot be null")
			return
		}
		fmt.Println("title: " + title)

		if password == "" {
			fmt.Println("password cannot be null")
			return
		}
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(quitmeetCmd)

	quitmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	quitmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
