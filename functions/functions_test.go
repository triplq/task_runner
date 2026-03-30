package functions

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	t.Run("Check", func(t *testing.T) {
		err := image_resize("../data/panda.jpg")
		fmt.Println(err)
	})
}
