package protocol

import "strings"

type Protocol int

const (
	Https Protocol = iota
	SSH
	Unknown
)

const (
	httpsStr = "https://"
	sshStr = "git@"
)

func GetProtocol(remoteUrl string) Protocol {
	if strings.HasPrefix(remoteUrl, httpsStr) {
		return Https
	}

	if strings.HasPrefix(remoteUrl, sshStr) {
		return SSH
	}

	return Unknown
}
