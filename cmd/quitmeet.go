package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
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
		 * 3.是否为会议的参与者
		 * 4.密码是否正确 
		 */

		curStatus := entity.CurStatus.GetStatus()		

		 // 1.是否登录
		has_login := curStatus.Islogin
		if !has_login {
			fmt.Println("GoAgenda quitmeet failed: You did not login yet!")
			return
		}

				
		// 2.会议是否存在
		meetingList := &service.MeetingModel
		exist := meetingList.IsExist(title)
		if !exist {
			fmt.Println("GoAgenda quitmeet failed: Meeting does not exist!")
			return
		}		
		
		// 3.是否为参与者		
		meeting := meetingList.GetMeeting(title)
		if !meeting.IsParticipator(curStatus.UserName) {
			fmt.Println("GoAgenda quitmeet failed: You are not a participator!")
			return
		}

		// 4.密码是否正确
		if password != curStatus.Password {
			fmt.Println("GoAgenda quitmeet failed: Invalid password!")
			return
		}

		meetingList.QuitMeeting(curStatus.UserName, title)
		fmt.Println("GoAgenda quitmeet succeed!")
	},
}

func init() {
	rootCmd.AddCommand(quitmeetCmd)

	quitmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	quitmeetCmd.Flags().StringP("password", "p", "", "用户密码")
}
