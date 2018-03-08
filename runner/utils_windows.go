package runner

import (
	"io"
	"os/exec"
	"strconv"
)

func killCmd(cmd *exec.Cmd) (int, error) {
	pid := cmd.Process.Pid
	// https://stackoverflow.com/a/44551450
	kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(pid))
	return pid, kill.Run()
}

func (e *Engine) startCmd(cmd string) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	var err error
	c := exec.Command("cmd", "/c", cmd)
	stderr, err := c.StderrPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return nil, nil, nil, err
	}
	err = c.Start()
	if err != nil {
		return nil, nil, nil, err
	}
	return c, stdout, stderr, err
}
