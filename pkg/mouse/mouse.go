package mouse

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
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
			robotgo.Click(Right)
		})
	},
}

var cmdDbRightClick = &cobra.Command{
	Use:   "db-riclick",
	Short: "双击鼠标右键",
	Run: func(cmd *cobra.Command, args []string) {
		clickWithInterval(interval, func() {
			robotgo.Click(Right, true)
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

// PressAndRelease 完成一次鼠标按下并释放的动作
func PressAndRelease(key string, interval time.Duration) error {
	t := time.NewTimer(interval)
	defer t.Stop()
	if err := robotgo.MouseDown(key); err != nil {
		return err
	}
	select {
	case <-t.C:
		if err := robotgo.MouseUp(key); err != nil {
			return err
		}
	}
	return nil
}
