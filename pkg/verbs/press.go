package verbs

import (
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	duration int
	unit     = time.Millisecond
)

func init() {
	CmdPress.AddCommand(resMouseLeft)
	CmdPress.AddCommand(resMouseRight)

	CmdPress.PersistentFlags().IntVarP(&duration, "duration", "d", -1, "按下时长\n<=0 表示不释放")
}

var CmdPress = &cobra.Command{
	Use:     "press",
	Aliases: []string{"pr"},
	Short:   "'按下'动作",
}

var resMouseLeft = &cobra.Command{
	Use:     "left-mouse",
	Aliases: []string{"lmouse", "lm"},
	Short:   "鼠标左键",
	RunE: func(cmd *cobra.Command, args []string) error {
		evtA := func() error {
			if err := robotgo.MouseDown(Left); err != nil {
				return err
			}
			return nil
		}
		evtB := func() error {
			if err := robotgo.MouseUp(Right); err != nil {
				return err
			}
			return nil
		}
		if err := evtWithDuration(evtA, evtB); err != nil {
			return err
		}
		return nil
	},
}

var resMouseRight = &cobra.Command{
	Use:     "right-mouse",
	Aliases: []string{"rmouse", "rm"},
	Short:   "鼠标右键",
	RunE: func(cmd *cobra.Command, args []string) error {
		a := func() error {
			if err := robotgo.MouseDown(Right); err != nil {
				return err
			}
			return nil
		}
		b := func() error {
			if err := robotgo.MouseUp(Right); err != nil {
				return err
			}
			return nil
		}
		if err := evtWithDuration(a, b); err != nil {
			return err
		}
		return nil
	},
}

func evtWithDuration(evtA, evtB func() error) error {
	if duration <= 0 {
		return evtA()
	}
	if err := evtA(); err != nil {
		return err
	}
	timer := time.NewTimer(time.Duration(duration) * unit)
	for {
		select {
		case <-timer.C:
			if err := evtB(); err != nil {
				return err
			}
			timer.Stop()
			return nil
		}
	}
}
