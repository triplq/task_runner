package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/triplq/task_runner/functions"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task",
	Long:  `Just add a task and it will apear in list`,
}

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resizing of image",
	Long: `U should enter a task name, img dest and (w,h) for scale it, then ull have resized image near original image
	Dont forger to check ur internet connection`,
	Run: func(cmd *cobra.Command, args []string) {
		w, err1 := strconv.Atoi(args[2])
		h, err2 := strconv.Atoi(args[3])
		if len(args) != 4 || err1 != nil || err2 != nil {
			fmt.Fprint(os.Stderr, "Wrong args, read description")
		} else {
			err := functions.Image_resize(args[1], w, h)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.AddCommand(resizeCmd)
}
