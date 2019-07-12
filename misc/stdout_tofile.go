package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
)

func main() {
	f, err := os.Create("./output.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cmd := exec.Command("hostname")

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(f)
	defer writer.Flush()

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()

}
