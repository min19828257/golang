package main

import (
	"fmt"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("cmd", "/C", "ipconfig")

	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> dates")
	fmt.Println(string(dateOut))
}
