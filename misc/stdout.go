package main

import (
	"os"
	"os/exec"
)

func main() {

	f, err := os.Create("./misc/out.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cmd := exec.Command("hostname")
	cmd.Stdout = f
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

}
