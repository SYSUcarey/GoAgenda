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
}
