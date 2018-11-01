package cmd

import (
	"fmt"
	"time"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
	"github.com/spf13/cobra"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "创建会议",
	Long: `
GoAgenda cm -t title -p participator -s starttime -e endtime

	各个参数分别对应:
	-t 会议标题
	-p 会议参与者
	-s 会议起始时间
	-e 会议结束时间`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("parti")
		starttime, _ := cmd.Flags().GetString("start")
		endtime, _ := cmd.Flags().GetString("end")

		/*
		 * 参数格式、逻辑处理
		 * 1. 登陆与否判断
		 * 2. 参数格式合法性判断
		 * 3. 参数逻辑合法性判断
		 */
		// 1.登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 未登陆的响应处理
		if !has_login {
			fmt.Println("GoAgenda cm failed: You did not login yet!")
			return
		}
		// 2. 参数格式合法性判断
		// 标题不能为空
		if title == "" {
			fmt.Println("GoAgenda cm failed: title cannot be null")
			return
		}
		// 参与者不能为空
		if participator == "" {
			fmt.Println("GoAgenda cm failed: participator cannot be null")
			return
		}
		// 起始时间不能为空
		if starttime == "" {
			fmt.Println("GoAgenda cm failed: starttime cannot be null")
			return
		}
		// 结束时间不能为空
		if endtime == "" {
			fmt.Println("GoAgenda cm failed: endtime cannot be null")
			return
		}
		// 起始结束时间必须合法（Format: "2006-01-02/15:04:05"）
		start, err_starttime_format_invalid := time.Parse("2006-01-02/15:04:05", starttime)
		if err_starttime_format_invalid != nil {
			fmt.Println("GoAgenda cm failed: starttime invalid!")
			fmt.Println("GoAgenda cm: starttime must be in format \"2006-01-02/15:04:05\"!")
			fmt.Println(err_starttime_format_invalid)
			return
		}
		end, err_endtime_format_invalid := time.Parse("2006-01-02/15:04:05", endtime)
		if err_endtime_format_invalid != nil {
			fmt.Println("GoAgenda cm failed: endtime invalid!")
			fmt.Println("GoAgenda cm: endtime must be in format \"2006-01-02/15:04:05\"!")
			fmt.Println(err_endtime_format_invalid)
			return
		}
		// 3.参数逻辑性判断
		// 会议标题不能重复
		if service.MeetingModel.IsExist(title) {
			fmt.Println("GoAgenda cm failed: title repeated!")
			return
		}
		// 参与者必须是一个用户
		if !service.UserModel.IsExist(participator) {
			fmt.Println("GoAgend cm failed: participator is not a user!")
			return
		}
		// 用户必须在会议时间有空
		// 读取meetings.json，遍历用户的所有会议时间，不能有冲突
		// todo
		// 参与者必须在会议时间有空
		// 读取meetings.json，遍历参与者的所有会议时间，不能有冲突
		// todo
		// 结束时间一定要在开始时间之后
		is_endtime_after_starttime := end.After(start)
		if !is_endtime_after_starttime {
			fmt.Println("GoAgenda cm failed: endtime must be after starttime!")
			return
		}
		// 开始时间一定要在当前时间之后
		currenttime := time.Now()
		is_starttime_after_currenttime := start.After(currenttime)
		if !is_starttime_after_currenttime {
			fmt.Println("GoAgenda cm failed: starttime must be after current time!")
			return
		}
	

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1. meetings.json添加一个会议
		 * 2. UI提示处理
		 */
		var participators = []string{}
		participators = append(participators, participator)
		if service.MeetingModel.CreateMeeting(entity.CurStatus.GetStatus().UserName, title, starttime, endtime, participators) {
			fmt.Println("GoAgenda cm succeed: ")
			fmt.Println("title: " + title)
			fmt.Println("participator: " + participator)
			fmt.Println("starttime: " + starttime)
			fmt.Println("endtime: " + endtime)
		} else {
			fmt.Println("GoAgenda cm failed: CreateMeeting failed!")
		}


	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "", "会议标题")
	cmCmd.Flags().StringP("parti", "p", "", "会议参与者")
	cmCmd.Flags().StringP("start", "s", "", "会议起始时间")
	cmCmd.Flags().StringP("end", "e", "", "会议结束时间")
}
