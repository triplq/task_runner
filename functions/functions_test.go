package functions

import (
	"fmt"
	"os"
	"testing"
)

func TestResize(t *testing.T) {
	t.Run("Is exist", func(t *testing.T) {
		file, err := os.Create("../data/file.txt")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		defer file.Close()

		err = os.Remove(file.Name())

		str, err := image_resize(file.Name())
		fmt.Println(str, err)
	})
}
