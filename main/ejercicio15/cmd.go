package cmd

import (
	"os/exec"
	"syscall"
)

func Power(cmd string) ([]byte, error) {
	c := exec.Command("powershell.exe", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := c.CombinedOutput()
	if err != nil {
		return nil, err

	}
	return output, err
}
