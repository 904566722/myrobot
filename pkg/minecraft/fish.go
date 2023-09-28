package minecraft

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	rbt "github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"

	"myrobot/pkg/mouse"
)

var (
	beforeFishColorPos string
	repeatNum          int
)

func init() {
	cmdFish.Flags().StringVarP(&beforeFishColorPos, "pos", "p", "", "判定鱼是否上钩会发生颜色变化的位置坐标")
	cmdFish.Flags().IntVarP(&repeatNum, "repeat", "r", 1, "执行钓鱼动作的次数")
}

var cmdFish = &cobra.Command{
	Use:     "auto-fish",
	Aliases: []string{"fish"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var x, y int
		if beforeFishColorPos == "" {
			x, y = rbt.Location()
		} else {
			var err error
			posArr := strings.Split(beforeFishColorPos, ",")
			if len(posArr) != 2 {
				return fmt.Errorf("坐标格式错误,当前:%s,应该形如:100,100", beforeFishColorPos)
			}
			x, err = strconv.Atoi(posArr[0])
			if err != nil {
				return err
			}
			y, err = strconv.Atoi(posArr[1])
			if err != nil {
				return err
			}
		}
		if err := autoFish(x, y, repeatNum); err != nil {
			return err
		}
		return nil
	},
}

// autoFish 自动钓鱼
// 启动流程之前，将鼠标放置于：与上钩之前与上钩之后会发生颜色变化的位置
func autoFish(x, y, repeatNum int) error {
	// 钓鱼中的标志颜色
	var colorFishing string
	// 抛竿
	if err := mouse.PressAndRelease(mouse.Left, 500*time.Millisecond); err != nil {
		return err
	}
	colorFishing = rbt.GetPixelColor(x, y)
	for repeatNum > 0 {
		if err := fishOnce(x, y, colorFishing); err != nil {
			return err
		}
		repeatNum--
	}
	return nil
}

// fishOnce 一次完整的钓鱼动作
// 1. 抛竿
// 2. 收杆
func fishOnce(x, y int, colorFishing string) error {
	// 抛竿
	if err := mouse.PressAndRelease(mouse.Left, 500*time.Millisecond); err != nil {
		return err
	}
	for {
		color := rbt.GetPixelColor(x, y)
		if color != colorFishing {
			// 颜色发生变化，鱼上钩
			rbt.Click()
			time.Sleep(100 * time.Millisecond)
			break
		}
	}
	return nil
}
