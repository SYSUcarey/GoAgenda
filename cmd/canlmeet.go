package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
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
			fmt.Println("title cannot be null")
			return
		}		
		fmt.Println("title: " , title)

		if password == "" {
			fmt.Println("password cannot be null")
			return
		}	
		fmt.Println("password: " , password)

		/*
		 * 合法性判断
		 * 1.是否登录
		 * 2.会议是否存在
		 * 3.密码是否正确	
		 */

		curStatus := entity.CurStatus.GetStatus()
		// 1.是否登录		
		
		// 2.会议是否存在

		// 3.密码是否正确
	},
}

func init() {
	rootCmd.AddCommand(canlmeetCmd)

	canlmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	canlmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
