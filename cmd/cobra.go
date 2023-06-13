package cmd

import (
	"fmt"
	"os"
	"tickor-cli/cmd/restart"
	"tickor-cli/cmd/start"
	"tickor-cli/cmd/stop"

	"github.com/spf13/cobra"
)

func tip() {
	usageStr := `欢迎使用 ` + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.go-admin.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

var rootCmd = &cobra.Command{
	Use:          "tickor-cli",
	Short:        "tickor-cli",
	SilenceUsage: true,
	Long:         `tickor-cli`,
	Example:      `awdw`,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) < 1 {
	// 		tip()
	// 		return errors.New("requires at least one arg")
	// 	}
	// 	return nil
	// },
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.SetHelpTemplate("daw")
	rootCmd.AddCommand(start.Cmd)
	rootCmd.AddCommand(stop.Cmd)
	rootCmd.AddCommand(restart.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
