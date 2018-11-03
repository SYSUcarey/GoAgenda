package cmd

import (
	"github.com/chenf99/GoAgenda/service"
	"github.com/chenf99/GoAgenda/entity"
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
		/*
		 * 参数格式、逻辑处理
		 * 1. 登陆与否判断
		 * 2. 参数格式合法性判断
		 * 3. 参数逻辑合法性判断
		 */

		
		// 1.登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 已经登陆无法进行注册命令
		if !has_login {
			service.Error.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  deluser failed: You did not login yet!")
			return
		}

		// 2. 参数格式合法性判断
		// 密码为空
		if	password == "" {
			service.Error.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  deluser failed: password cannot be null")
			return
		}

		// 3. 参数逻辑合法性判断
		// 当前用户名与密码参数是否匹配
		username := entity.CurStatus.GetStatus().UserName
		if !service.UserModel.MatchPass(username, password) {
			service.Error.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  deluser failed: username does not match password!")
			return
		}
		

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1.清空用户创建的会议 
		 * 2.退出用户参与的会议 
		 * 3.users.json delete当前用户
		 * 4.status.json 登出状态   
		 * 5.IO处理
		 */	

		curStatus := entity.CurStatus.GetStatus()
		// 1.清空用户创建的会议
		service.MeetingModel.EmptyMeeting(curStatus.UserName)
		
		// 2.退出用户参与的会议		
		// 获取所有会议
		meetingList := service.MeetingModel.Meetings
		for _, meeting := range meetingList {
			if meeting.IsParticipator(curStatus.UserName) {
				service.MeetingModel.RemoveMeetingParticipator(meeting.Sponsor, meeting.Title, curStatus.UserName)
			}
		}

		// 3.删除当前用户
		service.UserModel.DeleteUser(username)
		
		// 4.登出系统
		entity.CurStatus.LogOut()
		
		// 5.IO提示处理
		service.Info.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  deluser succeed!")

	},
}

func init() {
	rootCmd.AddCommand(deluserCmd)


	deluserCmd.Flags().StringP("pass", "p", "", "your password")
}
