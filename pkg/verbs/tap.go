package verbs

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var tapInterval int
var tapTimeUnit = time.Millisecond

var kbdBtn string

func init() {
	CmdTap.AddCommand(resLeftMouse)
	CmdTap.AddCommand(resKeyboard)

	CmdTap.PersistentFlags().IntVarP(&tapInterval, "interval", "i", -1, "点击间隔(毫秒)\n<=0 表示只点击一次")

	resKeyboard.Flags().StringVarP(&kbdBtn, "button", "b", "", "键盘按键(必需)")
	if err := resKeyboard.MarkFlagRequired("button"); err != nil {
		panic(err)
	}
}

var CmdTap = &cobra.Command{
	Use:   "tap",
	Short: "按下、点击动作",
}

var resLeftMouse = &cobra.Command{
	Use:     "left-mouse",
	Aliases: []string{"lmouse", "lm"},
	Short:   "鼠标左键",
	Run: func(cmd *cobra.Command, args []string) {
		tapWithInterval(func() {
			robotgo.Click()
		})
	},
}

func tapWithInterval(evt func()) {
	if tapInterval <= 0 {
		evt()
		return
	}

	ticker := time.NewTicker(time.Duration(tapInterval) * tapTimeUnit)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			evt()
		}
	}
}

func tapWithIntervalE(evt func() error) error {
	if tapInterval <= 0 {
		if err := evt(); err != nil {
			return err
		}
		return nil
	}
	ticker := time.NewTicker(time.Duration(tapInterval) * tapTimeUnit)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := evt(); err != nil {
				return err
			}
		}
	}
}

var resKeyboard = &cobra.Command{
	Use:     "keyboard",
	Aliases: []string{"kbd", "kb"},
	Short:   "键盘按键",
	RunE: func(cmd *cobra.Command, args []string) error {
		evt := func() error {
			if err := robotgo.KeyTap(kbdBtn); err != nil {
				return err
			}
			return nil
		}
		if err := tapWithIntervalE(evt); err != nil {
			return err
		}
		return nil
	},
}
