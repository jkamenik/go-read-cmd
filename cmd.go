package cmd

import (
	"io"
	"os/exec"
)

type Cmd struct {
	Name string
	Args []string

	cmd *exec.Cmd
}

func Command(name string, args ...string) *Cmd {
	cmd := new(Cmd)

	cmd.Name = name
	cmd.Args = args
	// cmd.Cmd = exec.Command(name, args...)

	return cmd
}

func (self *Cmd)Exec(input io.Reader) io.ReadCloser {
	cmd := exec.Command(self.Name, self.Args...)
	rpipe, wpipe := io.Pipe()

	cmd.Stdout = wpipe
	cmd.Stdin  = input

	go func() {
		err := cmd.Run()
		wpipe.CloseWithError(err)
	}()

	return rpipe
}
