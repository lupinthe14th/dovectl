package doveadm

import (
	"fmt"
	"os/exec"

	"github.com/cli/safeexec"
	"github.com/lupinthe14th/dovectl/models"
)

func doveadm() (string, error) {
	return safeexec.LookPath("doveadm")
}

func Backup(u *models.User) error {
	bin, err := doveadm()
	if err != nil {
		return err
	}
	if err := exec.Command(
		bin,
		"-o", fmt.Sprintf("imapc_user=%s", u.ID),
		"-o", fmt.Sprintf("imapc_password=%s", u.Password),
		"backup", "-R", "-u", u.ID, "imapc:",
	).Run(); err != nil {
		return fmt.Errorf("user: %v %v", u.ID, err)
	}
	return nil
}
func Sync(u *models.User) error {
	bin, err := doveadm()
	if err != nil {
		return err
	}
	if err := exec.Command(
		bin,
		"-o", fmt.Sprintf("imapc_user=%s", u.ID),
		"-o", fmt.Sprintf("imapc_password=%s", u.Password),
		"sync", "-R", "-u", u.ID, "imapc:",
	).Run(); err != nil {
		return fmt.Errorf("user: %v %v", u.ID, err)
	}
	return nil
}
