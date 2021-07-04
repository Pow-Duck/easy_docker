package easy_docker

import (
	"os/exec"
	"strings"
)

func command(cmd string) ([]string, error) {
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(output)), "\n"), err
}
