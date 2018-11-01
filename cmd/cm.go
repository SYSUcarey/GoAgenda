package cmd

import (
	"fmt"
	"time"
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
		has_login := true
		// 读取status.json判断是否已经登陆
		// todo
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
		// 结束时间一定要在开始时间之后
		is_endtime_after_starttime := end.After(start)
		if !is_endtime_after_starttime {
			fmt.Println("GoAgenda cm failed: endtime must be after starttime!")
			return
		}

		/*
		 * 参数格式、逻辑合法后的响应处理
		 */
		// todo

	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "", "会议标题")
	cmCmd.Flags().StringP("parti", "p", "", "会议参与者")
	cmCmd.Flags().StringP("start", "s", "", "会议起始时间")
	cmCmd.Flags().StringP("end", "e", "", "会议结束时间")
}
