package files

import (
	"fmt"
	"os/exec"
)


func OpenVSCode(path string) error {
	cmd := exec.Command("cmd", "/C", "code", path)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open VS Code: %v", err)
	}
	return nil
}

func OpenInWindows(path string) error {
	cmd := exec.Command("cmd", "/C", "start", "", path)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open: %v", err)
	}
	return nil
}
