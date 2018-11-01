package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
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

		/*
		 * 合法性判断
		 * 1.是否登录
		 * 2.会议是否存在
		 * 3.是否为会议的发起者
		 * 4.是否为会议的参与者
		 * 5.密码是否正确 
		 */

		curStatus := entity.CurStatus.GetStatus()		

		 // 1.是否登录
		has_login := curStatus.Islogin
		if !has_login {
			fmt.Println("GoAgenda cm failed: You did not login yet!")
			return
		}

		// 2.会议是否存在

		// 3.是否为发起者

		// 4.是否为参与者

		// 5.密码是否正确
		if password != curStatus.Password {
			fmt.Println("GoAgenda cm failed: Invalid password!")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(quitmeetCmd)

	quitmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	quitmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
