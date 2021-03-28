package doveadm

import (
	"fmt"
	"github.com/cli/safeexec"
	"github.com/lupinthe14th/sync/models"
	"os/exec"
)

func doveadm() (string, error) {
	return safeexec.LookPath("doveadm")
}

func Sync(u *models.User) error {
	bin, err := doveadm()
	if err != nil {
		return err
	}
	return exec.Command(
		bin,
		"-o", fmt.Sprintf("imapc_user=%s", u.ID),
		"-o", fmt.Sprintf("imapc_password=%s", u.Password),
		"sync", "-R", "-u", u.ID, "imapc:",
	).Run()
}