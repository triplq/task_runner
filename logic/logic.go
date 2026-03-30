package logic

import (
	"fmt"
	"os"
)

func image_resize(path string) (string, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("File is not exist")
	}

}

func worker(in <-chan int, out chan<- int) {
	task := <-in

}
