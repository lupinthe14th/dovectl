package usecases

import (
	"os/exec"

	"github.com/cli/safeexec"
)

func Echo(arg ...string) error {
	bin, err := safeexec.LookPath("echo")
	if err != nil {
		return err
	}
	return exec.Command(bin).Run()
}
