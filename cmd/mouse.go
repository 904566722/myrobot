package main

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"

	"myrobot/pkg/mouse"
)

var (
	interval int
)

func init() {
	CmdMouse.AddCommand(cmdClick)
	CmdMouse.AddCommand(cmdRightClick)
	CmdMouse.AddCommand(cmdDbClick)
	CmdMouse.AddCommand(cmdDbRightClick)

	CmdMouse.PersistentFlags().IntVarP(&interval, "interval", "i", -1, "事件的时间间隔(毫秒)\n<=0 表示只点击一次")

	//cmdClick.Flags().IntVarP(&interval, "interval", "i", -1, "每次点击的事件间隔(毫秒)\n<=0 表示只点击一次")
}

var CmdMouse = &cobra.Command{
	Use:   "mouse",
	Short: "鼠标相关事件",
}

var cmdClick = &cobra.Command{
	Use:   "click",
	Short: "点击鼠标左键",
	Run: func(cmd *cobra.Command, args []string) {
		clickWithInterval(interval, func() {
			robotgo.Click(robotgo.Left)
		})
	},
}

var cmdDbClick = &cobra.Command{
	Use:   "db-click",
	Short: "双击鼠标左键",
	Run: func(cmd *cobra.Command, args []string) {
		clickWithInterval(interval, func() {
			robotgo.Click(robotgo.Left, true)
		})
	},
}

var cmdRightClick = &cobra.Command{
	Use:   "riclick",
	Short: "点击鼠标右键",
	Run: func(cmd *cobra.Command, args []string) {
		clickWithInterval(interval, func() {
			robotgo.Click(mouse.Right)
		})
	},
}

var cmdDbRightClick = &cobra.Command{
	Use:   "db-riclick",
	Short: "双击鼠标右键",
	Run: func(cmd *cobra.Command, args []string) {
		clickWithInterval(interval, func() {
			robotgo.Click(mouse.Right, true)
		})
	},
}

func clickWithInterval(interval int, clickEvt func()) {
	if interval <= 0 {
		clickEvt()
		return
	}

	for {
		clickEvt()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
