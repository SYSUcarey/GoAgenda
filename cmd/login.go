package cmd

import (
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
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
		/*
		 * 参数格式、逻辑处理
		 * 1. 登陆与否判断
		 * 2. 参数格式合法性判断
		 * 3. 参数逻辑合法性判断
		 */
		// 1.登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 已经登陆无法进行注册命令
		if has_login {
			service.Error.Println("GoAgenda " + username + "  login failed: You had already logined!")
			return
		}

		// 2. 参数格式合法性判断
		// 用户名为空
		if	username == "" {
			service.Error.Println("GoAgenda " + username + "  login failed: username cannot be null")
			return
		}
		// 密码为空
		if	password == "" {
			service.Error.Println("GoAgenda " + username + "  login failed: password cannot be null")
			return
		}

		// 3. 参数逻辑合法性判断
		// 登陆用户名必须存在
		if !service.UserModel.IsExist(username) {
			service.Error.Println("GoAgenda " + username + "  login failed: username does not exist!")
			return
		}
		// 用户名和密码必须匹配
		if !service.UserModel.MatchPass(username, password) {
			service.Error.Println("GoAgenda " + username + "  login failed: username does not match password!")
			return
		}
		

		/*
		 * 参数格式、逻辑合法后的响应处理
		 * 1. status.json添加登陆状态
		 * 2. 写入日志并UI提示
		 */	
		
		entity.CurStatus.LogIn(username, password)
		service.Info.Println("GoAgenda " + username + "  login succeed!")


	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("user", "u", "", "your username")
	loginCmd.Flags().StringP("pass", "p", "", "your password")
}
