package keyboard

import (
	"errors"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	interval int
)

var (
	ErrNotEqOneLetter = errors.New("仅允许传入一个字母")
	ErrKeyTap         = errors.New("keyTap失败")
)

func init() {
	CmdKeyBoard.AddCommand(cmdTapLetter)
	CmdKeyBoard.AddCommand(cmdPress)
	CmdKeyBoard.AddCommand(cmdRelease)

	CmdKeyBoard.PersistentFlags().IntVarP(&interval, "interval", "i", -1, "事件的时间间隔(毫秒)\n<=0 表示只键入一次")
}

var CmdKeyBoard = &cobra.Command{
	Use:   "key",
	Short: "键盘相关事件",
}

var cmdTapLetter = &cobra.Command{
	Use:   "tap-letter",
	Short: "键入字母",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || len(args[0]) != 1 {
			return ErrNotEqOneLetter
		}
		return keyTapWithInterval(interval, func() error {
			if err := robotgo.KeyTap(args[0]); err != nil {
				return err
			}
			return nil
		})
	},
}

var cmdPress = &cobra.Command{
	Use:   "press",
	Short: "按下某个按键",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || len(args[0]) != 1 {
			return ErrNotEqOneLetter
		}
		return robotgo.KeyDown(args[0])
	},
}

var cmdRelease = &cobra.Command{
	Use:   "release",
	Short: "释放某个按键",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || len(args[0]) != 1 {
			return ErrNotEqOneLetter
		}
		return robotgo.KeyUp(args[0])
	},
}

func keyTapWithInterval(interval int, evt func() error) error {
	if interval <= 0 {
		return evt()
	}

	for {
		if err := evt(); err != nil {
			return err
		}
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
