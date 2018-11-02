package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
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
		
		/*
		 * 合法性判断
		 * 1.是否登录
		 * 2.密码是否正确  
		 */ 

		curStatus := entity.CurStatus.GetStatus()

		// 1.是否登录
		has_login := curStatus.Islogin
		if !has_login {
			fmt.Println("GoAgenda emptymeet failed: You did not login yet!")
			return
		}

		// 2.密码是否正确
		if password != curStatus.Password {
			fmt.Println("GoAgenda emptymeet failed: Invalid password!")
			return
		}

		// 参数合法
		fmt.Println("GoAgenda emptymeet succeed!")
		service.MeetingModel.EmptyMeeting(curStatus.UserName)
	},
}

func init() {
	rootCmd.AddCommand(emptymeetCmd)

	emptymeetCmd.Flags().StringP("password", "p", "", "your password")
}
