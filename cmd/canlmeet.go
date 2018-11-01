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

// canlmeetCmd represents the canlmeet command
var canlmeetCmd = &cobra.Command{
	Use:   "canlmeet",
	Short: "取消会议",
	Long: `
GoAgenda canlmeet -t title -p password
	
	各个参数分别对应:
	-t 会议标题
	-p 用户密码`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		password, _ := cmd.Flags().GetString("password")
		//处理参数
		if title == "" {
			fmt.Println("title cannot be null" + title)
			return
		}		
		fmt.Println("title: " + title)

		if password == "" {
			fmt.Println("password cannot be null" + title)
			return
		}	
		fmt.Println("password: " + password)
	},
}

func init() {
	rootCmd.AddCommand(canlmeetCmd)

	canlmeetCmd.Flags().StringP("title", "t", "", "会议标题")
	canlmeetCmd.Flags().StringP("password", "p", "", "用户密码")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// canlmeetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// canlmeetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
