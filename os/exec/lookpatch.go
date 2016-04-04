package main

import (
	"fmt"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ls is available at %s\n", path)
	}

	path, err = exec.LookPath("hello")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("hello is available at %s\n", path)
	}

}
