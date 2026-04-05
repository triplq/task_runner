package functions

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	t.Run("Check", func(t *testing.T) {
		err := Image_resize("../data/panda.jpg", 600, 600)
		fmt.Println(err)
	})
}
