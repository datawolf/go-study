package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(out))
}
