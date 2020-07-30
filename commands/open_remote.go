package commands

import (
	"github.com/cyberfox1002/groo/replacer"
	"os/exec"
)

func OpenRemote(origin string) error {
	out, err := exec.Command("git", "remote", "get-url", origin).Output()
	if err != nil {
		return err
	}

	url, err := replacer.GetRemoteSiteUrl(string(out))
	if err != nil {
		return err
	}

	err = exec.Command("open", url).Run()
	if err != nil {
		return err
	}

	return nil
}