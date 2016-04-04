package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	stdout, err := cmd.StdoutPipe()
	cmd.Start()

	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(content))
}
