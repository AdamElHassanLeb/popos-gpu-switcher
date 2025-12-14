package service

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func run(cmd string, args ...string) (string, error) {
	c := exec.Command(cmd, args...)

	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run() // blocks until finished
	if err != nil {
		if stderr.Len() > 0 {
			return "", errors.New(strings.TrimSpace(stderr.String()))
		}
		return "", err
	}

	return strings.TrimSpace(stdout.String()), nil
}
