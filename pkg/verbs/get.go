package verbs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	colorPos string
)

func init() {
	CmdGet.AddCommand(cmdColor)
	CmdGet.AddCommand(cmdPos)

	cmdColor.Flags().StringVarP(&colorPos, "position", "p", "", "位置，形如'100,100'")
}

var CmdGet = &cobra.Command{
	Use:   "get",
	Short: "'获取'动作",
}

var cmdColor = &cobra.Command{
	Use:     "color",
	Short:   "获取颜色",
	Long:    `获取对应位置的颜色，位置由两位数坐标表示，形如 1 2，如果不传入位置信息，则获取鼠标当前位置的颜色`,
	Aliases: []string{"cl"},
	RunE: func(cmd *cobra.Command, args []string) error {
		// 参数校验:只能传入 x y 坐标，或者不传(使用鼠标当前的坐标)
		var x, y int
		if colorPos == "" {
			x, y = robotgo.Location()
		} else {
			posArr := strings.Split(colorPos, ",")
			var err error
			x, err = strconv.Atoi(posArr[0])
			if err != nil {
				return err
			}
			y, err = strconv.Atoi(posArr[1])
			if err != nil {
				return err
			}
		}
		color := robotgo.GetPixelColor(x, y)
		fmt.Println(color)
		return nil
	},
}

var cmdPos = &cobra.Command{
	Use:     "position",
	Aliases: []string{"pos"},
	Short:   "获取鼠标当前位置",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(robotgo.Location())
	},
}
