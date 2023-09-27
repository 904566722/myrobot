package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	time.Sleep(2 * time.Second)
	fmt.Printf("获取鼠标位置中\n")
	x, y := robotgo.Location()
	fmt.Printf("鼠标位置:(%d,%d)\n", x, y)
}
