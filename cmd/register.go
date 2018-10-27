// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		//fmt.Println("register called")
		cmd.Help()
		username, _ := cmd.Flags().GetString("user")
		fmt.Println("register called by " + username)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("user", "u", "Anonymous", "your username")
	registerCmd.Flags().StringP("pass", "p", "", "your password")
	registerCmd.Flags().StringP("email", "e", "", "your email URL")
	registerCmd.Flags().StringP("tel", "t", "", "your telephone number")
}
