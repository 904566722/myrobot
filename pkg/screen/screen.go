package screen

import "github.com/spf13/cobra"

func init() {
	CmdScreen.AddCommand()
}

var CmdScreen = &cobra.Command{
	Use:     "screen",
	Aliases: []string{"sc"},
	Short:   "屏幕相关事件",
}
