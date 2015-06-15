package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	cmd := os.Args[1]
	args := os.Args[2:]
	err := do(cmd, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func do(run string, args ...string) error {
	cmd := exec.Command(run, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	defer stderr.Close()

	go read(stderr, os.Stderr, "stderr")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	defer stdout.Close()

	go read(stdout, os.Stdout, "stdout")

	err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

var regex = regexp.MustCompile("^(.+):(\\d+)")

func read(r io.Reader, w io.Writer, prefix string) {
	wd, _ := os.Getwd()

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if regex.Match([]byte(line)) {
			line = wd + "/" + line
		}
		fmt.Fprintf(w, "%s\n", line)
	}
	if s.Err() != nil {
		fmt.Printf("got err: %q", s.Err())
	}
}
