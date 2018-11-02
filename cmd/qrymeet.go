package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"time"
	"github.com/chenf99/GoAgenda/service"
	"github.com/spf13/cobra"
)

// qrymeetCmd represents the qrymeet command
var qrymeetCmd = &cobra.Command{
	Use:   "qrymeet",
	Short: "查询会议",
	Long: `
GoAgenda qrymeet -s starttime -e endtime

	各个参数分别对应:
	-s 会议起始时间
	-e 会议结束时间`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		starttime, _ := cmd.Flags().GetString("start")
		endtime, _ := cmd.Flags().GetString("end")
		//处理参数		
		if starttime == "" {
			fmt.Println("starttime cannot be null")
			return
		}
		fmt.Println("starttime: " + starttime)

		if endtime == "" {
			fmt.Println("endtime cannot be null")
			return
		}
		fmt.Println("endtime: " + endtime)

		/*
		 * 合法性判断
		 * 1.是否登录
		 * 2.时间格式是否合法
		 * 3.开始时间是否小于结束时间   
		 * 4.开始时间大于当前时间 
		 */		 

		curStatus := entity.CurStatus.GetStatus()

		// 1.是否登录
		has_login := curStatus.Islogin
		if !has_login {
			fmt.Println("GoAgenda qrymeet failed: You did not login yet!")
			return
		}

		// 2.时间格式是否合法
		start, err_starttime_format_invalid := time.Parse("2006-01-02/15:04:05", starttime)
		if err_starttime_format_invalid != nil {
			fmt.Println("GoAgenda qrymeet failed: starttime invalid!")
			fmt.Println("GoAgenda qrymeet: starttime must be in format \"2006-01-02/15:04:05\"!")
			fmt.Println(err_starttime_format_invalid)
			return
		}

		end, err_endtime_format_invalid := time.Parse("2006-01-02/15:04:05", endtime)
		if err_endtime_format_invalid != nil {
			fmt.Println("GoAgenda qrymeet failed: endtime invalid!")
			fmt.Println("GoAgenda qrymeet: endtime must be in format \"2006-01-02/15:04:05\"!")
			fmt.Println(err_endtime_format_invalid)
			return
		}

		// 3.开始时间是否小于结束时间
		is_endtime_after_starttime := end.After(start)
		if !is_endtime_after_starttime {
			fmt.Println("GoAgenda qrymeet failed: endtime must be after starttime!")
			return
		}

		// 4.开始时间大于当前时间
		currenttime := time.Now()
		is_starttime_after_currenttime := start.After(currenttime)
		if !is_starttime_after_currenttime {
			fmt.Println("GoAgenda qrymeet failed: starttime must be after current time!")
			return
		}

		// 参数合法
		fmt.Println("GoAgenda qrymeet succeed")
		meetingList := service.MeetingModel.MeetingQuery(curStatus.UserName, starttime, endtime)
		fmt.Println("There are ", len(meetingList), "meeting you sponsor or participate:")
		for i,meeting := range meetingList {
			fmt.Println("Meeting ", i + 1)
			fmt.Println("Title: ",meeting.GetTitle())
			fmt.Println("Sponsor: ",meeting.GetSponsor())
			fmt.Println("Participators: ",meeting.GetParticipator())
			fmt.Println("Start: ",meeting.GetStartDate())
			fmt.Println("End: ",meeting.GetEndDate())
		}
	},
}

func init() {
	rootCmd.AddCommand(qrymeetCmd)

	qrymeetCmd.Flags().StringP("start", "s", "", "会议起始时间")
	qrymeetCmd.Flags().StringP("end", "e", "", "会议结束时间")
}
