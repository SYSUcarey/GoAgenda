package cmd

import (
	"fmt"
	"github.com/chenf99/GoAgenda/entity"
	"github.com/spf13/cobra"
)

// qryuserCmd represents the qryuser command
var qryuserCmd = &cobra.Command{
	Use:   "qryuser",
	Short: "用户查询",
	Long: `
GoAgenda qryuser
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// 登陆与否判断
		has_login := entity.CurStatus.GetStatus().Islogin
		// 读取status.json判断是否已经登陆
		// todo
		// 已经登陆无法进行注册命令
		if !has_login {
			fmt.Println("GoAgenda qryuser failed: You did not login yet!")
			return
		}

		/*
		 * 状态、参数格式、逻辑合法后的响应处理
		 * 1. 读取users.json，获得用户列表
		 * todo
		 */	
		fmt.Println("GoAgenda qryuser succeed: ")
		user_list := entity.UserModel.GetAllUsers()
		fmt.Println("There are ", len(user_list), " users!")
		for _, user := range user_list {
			fmt.Println(user.Name)
		}

	},
}

func init() {
	rootCmd.AddCommand(qryuserCmd)
}
