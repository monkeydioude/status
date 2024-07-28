package service

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

func GetServiceStatus(serviceName string) (string, error) {
	cmd := exec.Command("systemctl", "status", serviceName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Filter out lines containing `CGroup:`
	var result strings.Builder
	scanner := bufio.NewScanner(&out)
	cGroupFlag := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "CGroup:") {
			cGroupFlag = true
		}
		if line == "\n" || line == "" {
			cGroupFlag = false
		}
		if !cGroupFlag {
			result.WriteString(line + "\n")
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return result.String(), nil
}
