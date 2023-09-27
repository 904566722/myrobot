package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"myrobot/pkg/keyboard"
	"myrobot/pkg/mouse"
)

// 好的命令行程序应该遵守: APPNAME VERB NOUN --ADJECTIVE 或 APPNAME COMMAND ARG --FLAG
// e.g.
// ./myrobot mouse click --interval=1000

var rootCmd = &cobra.Command{
	Use: "myrobot",
}

func main() {
	rootCmd.AddCommand(mouse.CmdMouse)       // 添加鼠标事件
	rootCmd.AddCommand(keyboard.CmdKeyBoard) // 添加键盘事件

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
