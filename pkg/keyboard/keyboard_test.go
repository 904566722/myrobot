package keyboard

import (
	"sync"
	"testing"
	"time"

	"github.com/go-vgo/robotgo"
)

func Test_keyTap(t *testing.T) {
	if err := robotgo.KeyTap("a"); err != nil {
		t.Error(err)
		return
	}
}

func Test_PressRelease(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := robotgo.KeyDown("a"); err != nil {
			t.Error(err)
			return
		}
		wg.Done()
	}()
	wg.Wait()
	time.Sleep(2 * time.Second)
}

func Test_keyTapWithInterval(t *testing.T) {
	if err := keyTapWithInterval(-1, func() error {
		if err := robotgo.KeyTap("a"); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Error(err)
		return
	}

}
