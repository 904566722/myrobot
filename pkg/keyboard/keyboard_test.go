package keyboard

import (
	"testing"

	"github.com/go-vgo/robotgo"
)

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
