package utils

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// PosChange 识别到屏幕某个点发生变化
func PosChange(x, y int, beforeColor string) {
	for {
		afterColor := robotgo.GetPixelColor(x, y)
		if afterColor != beforeColor {
			fmt.Printf("(%d,%d)坐标位置的颜色发生变化(%s-->%s)", x, y, beforeColor, afterColor)
			break
		}
	}
}
