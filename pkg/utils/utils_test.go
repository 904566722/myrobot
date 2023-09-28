package utils

import (
	"testing"

	"github.com/go-vgo/robotgo"
)

func TestPosChange(t *testing.T) {
	x, y := 1807, 537
	beforColor := robotgo.GetPixelColor(x, y)
	PosChange(x, y, beforColor)
}
