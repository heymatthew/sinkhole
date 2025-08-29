package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Problem finding home:" + err.Error())
	}
	file, err := os.Create(home + "/log/sinkhole.log")

	if err != nil {
		panic("Problem creating file:" + err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, os.Stdin)
	if err != nil {
		panic("Error copying data:" + err.Error())
	}
	fmt.Println("Fin")
}
