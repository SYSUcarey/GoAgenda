package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
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
		 * 3.是否有权删除		
		 * 4.密码是否正确	
		 */

		curStatus := entity.CurStatus.GetStatus()
		// 1.是否登录		
		has_login := curStatus.Islogin
		if !has_login {
			fmt.Println("GoAgenda canlmeet failed: You did not login yet!")
			return
		}
				
		meetingList := service.MeetingModel
		// 2.会议是否存在
		exist := (&meetingList).IsExist(title)
		if !exist {
			fmt.Println("GoAgenda canlmeet failed: Meeting does not exist!")
			return
		}
		
		// 3.是否有权删除（是否为发起者）
		meeting := (&meetingList).GetMeeting(title)
		user := curStatus.UserName
		if user != meeting.GetSponsor() {
			fmt.Println("GoAgenda canlmeet failed: You are not the sponsor!")
			return
		}

		// 4.密码是否正确
		if password != curStatus.Password {
			fmt.Println("GoAgenda canlmeet failed: Invalid password!")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(canlmeetCmd)

	canlmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	canlmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
