package mouse

import (
	"fmt"
	"testing"
	"time"
)

func TestPressAndRelease2(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("aaa")
	for {
		select {
		case <-timer.C:
			fmt.Println("bbb")
		}
	}
}
