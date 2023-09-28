package minecraft

import "github.com/spf13/cobra"

func init() {
	CmdMinecraft.AddCommand(cmdFish)
}

var CmdMinecraft = &cobra.Command{
	Use:     "minecraft",
	Aliases: []string{"mc"},
	Short:   "我的世界自动化命令",
}
