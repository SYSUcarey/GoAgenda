package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"

	"github.com/spf13/cobra"
)

// delparCmd represents the delpar command
var delparCmd = &cobra.Command{
	Use:   "delpar",
	Short: "删除会议参与者",
	Long: `
GoAgenda delpar -t title -p participator
	
	各个参数分别对应:
	-t 会议标题
	-p 会议参与者
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("parti")
		//处理参数

		// 参数逻辑合法性判断

		if title == "" {
			fmt.Println("GoAgenda delpar failed: title cannot be null")
			return
		}
		fmt.Println("title: " + title)

		if participator == "" {
			fmt.Println("GoAgenda delpar failed: participator cannot be null")
			return
		}

		if !(service.MeetingModel.GetMeeting(title).GetSponsor() == entity.CurStatus.UserName) {
			fmt.Println("GoAgenda delpar failed: you must be the sponsor of the meeting")
			return
		} // 当前用户必须是该会议的发起者

		if !service.MeetingModel.IsExist(title) {
			fmt.Println("GoAgenda delpar failed: this string is not one of meeting's title")
			return
		} // 必须存在有这个标题的会议

		if !service.MeetingModel.GetMeeting(title).IsParticipator(participator) {
			fmt.Println("GoAgenda delpar failed: participator is not in the meeting")
			return
		} // 会议中有这个参与者

		//service.MeetingModel.GetMeeting(title).RemoveParticipator(participator)

		fmt.Println("participator: " + participator)
	},
}

func init() {
	rootCmd.AddCommand(delparCmd)
	delparCmd.Flags().StringP("title", "t", "", "会议标题")
	delparCmd.Flags().StringP("parti", "p", "", "会议参与者")
}
