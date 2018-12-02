package process

import (
	"io"
	"io/ioutil"
	"os/exec"
)

// NewProcessWithMiddleModifier spawn process wrapper
func NewProcessWithMiddleModifier(name string, params []string, midProcess func(*exec.Cmd)) (Killer, error) {
	cmd := exec.Command(name, params...)
	machineSubProcess(cmd)
	midProcess(cmd)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd.Process, nil
}

// NewProcess will spawn new process
func NewProcess(name string, params ...string) (Killer, error) {
	return NewProcessWithMiddleModifier(name, params, func(cmd *exec.Cmd) {
	})
}

// NewProcessStdout will spawn new process and listen its Stdout
func NewProcessStdout(name string, params ...string) (stdout io.ReadCloser, killer Killer, err error) {
	killer, err = NewProcessWithMiddleModifier(name, params, func(cmd *exec.Cmd) {
		stdout, _ = cmd.StdoutPipe()
	})
	return
}

// NewProcessAllStdout spawn process and real all stdout
func NewProcessAllStdout(name string, params ...string) ([]byte, error) {
	stdout, _, err := NewProcessStdout(name, params...)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(stdout)
}
