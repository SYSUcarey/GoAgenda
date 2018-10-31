package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "用户注册",
	Long: `
GoAgenda register -u username -p password -e email -t telephone 
各个参数分别对应用户名、密码、邮箱地址、电话号码`,
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数值
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("pass")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("tel")
		// 处理参数
		if username == "" {
			fmt.Println("username cannot be null")
			return
		} else {			
			fmt.Println("username: " + username)
		}
		if password == "" {
			fmt.Println("password cannot be null")
			return
		} else {
			fmt.Println("password: " + password)
		}
		if email == "" {
			fmt.Println("email cannot be null")
			return
		} else {
			fmt.Println("email: " + email)
		}
		if telephone == "" {
			fmt.Println("telephone cannot be null")
			return
		} else {
			fmt.Println("telephone: " + telephone)
		}					
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringP("user", "u", "", "your username")
	registerCmd.Flags().StringP("pass", "p", "", "your password")
	registerCmd.Flags().StringP("email", "e", "", "your email URL")
	registerCmd.Flags().StringP("tel", "t", "", "your telephone number")
}
