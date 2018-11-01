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

// qrymeetCmd represents the qrymeet command
var qrymeetCmd = &cobra.Command{
	Use:   "qrymeet",
	Short: "查询会议",
	Long: `
GoAgenda qrymeet -s starttime -e endtime

	各个参数分别对应:
	-s 会议起始时间
	-e 会议结束时间`,
	Run: func(cmd *cobra.Command, args []string) {
		//获取参数
		starttime, _ := cmd.Flags().GetString("start")
		endtime, _ := cmd.Flags().GetString("end")
		//处理参数		
		if starttime == "" {
			fmt.Println("starttime cannot be null")
			return
		}
		fmt.Println("starttime: " + starttime)

		if endtime == "" {
			fmt.Println("endtime cannot be null")
			return
		}
		fmt.Println("endtime: " + endtime)
	},
}

func init() {
	rootCmd.AddCommand(qrymeetCmd)

	qrymeetCmd.Flags().StringP("start", "s", "", "会议起始时间")
	qrymeetCmd.Flags().StringP("end", "e", "", "会议结束时间")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// qrymeetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// qrymeetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
