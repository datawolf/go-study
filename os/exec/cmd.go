package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := &exec.Cmd{
		Path: "/home/datawolf/work/bin/hello",
		Args: append([]string{"ls"}, "-l"),
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}
