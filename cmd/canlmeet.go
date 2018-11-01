package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// canlmeetCmd represents the canlmeet command
var canlmeetCmd = &cobra.Command{
	Use:   "canlmeet",
	Short: "取消会议",
	Long: `
GoAgenda canlmeet -t title -p password
	
	各个参数分别对应:
	-t 会议标题
	-p 用户密码`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		password, _ := cmd.Flags().GetString("password")
		//处理参数
		if title == "" {
			fmt.Println("title cannot be null" + title)
			return
		}		
		fmt.Println("title: " + title)

		if password == "" {
			fmt.Println("password cannot be null" + title)
			return
		}	
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(canlmeetCmd)

	canlmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	canlmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
