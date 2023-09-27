package mouse

import "github.com/go-vgo/robotgo"

func Click() {

}

func LeftClick() {
	robotgo.Click()
}

func RightClick() {
	robotgo.Click(Right)
}

func DbLeftClick() {
	robotgo.Click(Left, true)
}

func DbRightClick() {
	robotgo.Click(Right, true)
}
