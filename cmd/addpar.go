package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addparCmd represents the addpar command
var addparCmd = &cobra.Command{
	Use:   "addpar",
	Short: "增加会议参与者",x
	Long: `
GoAgenda addpar -t title -p participator

	各个参数分别对应:
	-t 会议标题
	-p 会议参与者`,
	Run: func(cmd *cobra.Command, args []string) {		
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("parti")
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
			fmt.Println("GoAgenda addpar failed: You did not login yet!")
			return
		}
		// 2. 参数格式合法性判断
		// 标题不能为空
		if title == "" {
			fmt.Println("GoAgenda addpar failed: title cannot be null")
			return
		}
		// 添加的参与者不能为空
		if participator == "" {
			fmt.Println("GoAgenda addpar failed: participator cannot be null")
			return
		}
		// 3.参数逻辑合法性判断
		// 必须存在有这个标题的会议 
		// todo
		// 当前用户必须是该会议的发起者
		// todo
		// 必须存在参与者这一个用户
		// todo
		// 会议中原本没有这个参与者
		// todo
		// 参与者必须有时间参加会议
		// todo

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1. meetings.json对应会议中添加一个participator
		 * 
		 */
		// todo
		fmt.Println("GoAgenda addpar succeed: ")
		fmt.Println("title: ", title)
		fmt.Println("participator: ", participator)
	},
}

func init() {
	rootCmd.AddCommand(addparCmd)

	addparCmd.Flags().StringP("title", "t", "", "会议标题")
	addparCmd.Flags().StringP("parti", "p", "", "会议参与者")
}