package libwimark

import (
	"bytes"
	"os/exec"
)

func ExecCommand(name string, params ...string) (string, error) {

	cmd := exec.Command(name, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
