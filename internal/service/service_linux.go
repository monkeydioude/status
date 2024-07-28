package service

import "os/exec"

func GetServiceStatus(serviceName string) (string, error) {
	cmd := exec.Command("systemctl", "status", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	outputStr := string(output)
	return outputStr, nil
}
