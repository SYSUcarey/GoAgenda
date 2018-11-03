package cmd

import (
	"github.com/chenf99/GoAgenda/entity"
	"github.com/chenf99/GoAgenda/service"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "用户登出",
	Long: `
GoAgenda logout
`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// 登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 已经登陆无法进行注册命令
		if !has_login {
			service.Error.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  logout failed: You did not login yet!")
			return
		}

		/*
		 * 状态、参数格式、逻辑合法后的响应处理
		 * 1. status.json登出状态
		 * 2. 写入日志并UI提示
		 */	
		service.Info.Println("GoAgenda " + entity.CurStatus.GetStatus().UserName + "  logout succeed! ")
		entity.CurStatus.LogOut()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
