package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// delparCmd represents the delpar command
var delparCmd = &cobra.Command{
	Use:   "delpar",
	Short: "删除会议参与者",
	Long: `
GoAgenda delpar -t title -p participator
	
	各个参数分别对应:
	-t 会议标题
	-p 会议参与者
	`,
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
	rootCmd.AddCommand(delparCmd)
	delparCmd.Flags().StringP("title", "t", "", "会议标题")
	delparCmd.Flags().StringP("parti", "p", "", "会议参与者")
}
