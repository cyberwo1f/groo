package commands

import (
	"bytes"
	"errors"
	"github.com/cyberfox1002/groo/replacer"
	"os/exec"
)

func OpenRemote(origin string) error {
	var out string
	var err error

	out, err = exeCmd("git", "remote", "get-url", origin)
	if err != nil {
		return err
	}

	url, err := replacer.GetRemoteSiteUrl(out)
	if err != nil {
		return err
	}

	out, err = exeCmd("open", url)
	if err != nil {
		return err
	}

	return nil
}

func exeCmd(name string, arg ...string) (string, error) {
	var cmd *exec.Cmd
	var out bytes.Buffer
	var stderr bytes.Buffer
	var err error

	cmd = exec.Command(name, arg...)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return out.String(), errors.New(stderr.String())
	}

	return out.String(), nil
}
