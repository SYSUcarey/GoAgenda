package cmd

import (
	"fmt"

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
		has_login := true
		// 读取status.json判断是否已经登陆
		// todo
		// 已经登陆无法进行注册命令
		if !has_login {
			fmt.Println("GoAgenda deluser failed: You did not login yet!")
			return
		}

		// 2. 参数格式合法性判断
		// 密码为空
		if	password == "" {
			fmt.Println("GoAgenda deluser failed: password cannot be null")
			return
		}

		// 3. 参数逻辑合法性判断
		// 当前用户名与密码参数是否匹配
		// todo
		

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1. status.json添加一个用户
		 * 
		 */	
		fmt.Println("GoAgenda deluser succeed: ")
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(deluserCmd)


	deluserCmd.Flags().StringP("pass", "p", "", "your password")
}
