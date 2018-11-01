package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addparCmd represents the addpar command
var addparCmd = &cobra.Command{
	Use:   "addpar",
	Short: "增加会议参与者",
	Long: `
GoAgenda addpar -t title -p participator

	各个参数分别对应:
	-t 会议标题
	-p 会议参与者`,
	Run: func(cmd *cobra.Command, args []string) {		
		//获取参数
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("parti")
		//处理参数
		if title == "" {
			fmt.Println("title cannot be null")
			return
		}
		fmt.Println("title: " + title)

		if participator == "" {
			fmt.Println("participator cannot be null")
			return
		}
		fmt.Println("participator: " + participator)
	},
}

func init() {
	rootCmd.AddCommand(addparCmd)

	addparCmd.Flags().StringP("title", "t", "", "会议标题")
	addparCmd.Flags().StringP("parti", "p", "", "会议参与者")
}
