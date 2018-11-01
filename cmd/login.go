package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "用户登陆",
	Long: 
	`
GoAgenda login -u username -p password

	各个参数分别对应:
	-u 用户名
	-p 用户密码
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
		// 处理参数
		// 用户名为空
		if	username == "" {
			fmt.Println("username cannot be null")
			return
		}
		fmt.Println("username: " + username)

		// 密码为空
		if	password == "" {
			fmt.Println("password cannot be null")
			return
		}
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "your username")
	loginCmd.Flags().StringP("pass", "p", "", "your password")
}
