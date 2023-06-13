package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "这是删除的简短介绍",
	Long:  `这是删除的长介绍.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("我是执行删除操作的!")
	},
}

var All = &cobra.Command{
	Use:   "all",
	Short: "这是删除的简短介绍",
	Long:  `这是删除的长介绍.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete all!")
	},
}

func init() {
	Cmd.AddCommand(All)
}
