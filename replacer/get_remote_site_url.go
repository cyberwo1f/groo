package replacer

import (
	"github.com/cyberfox1002/groo/domain"
	"github.com/cyberfox1002/groo/protocol"
	"strings"
)

/*
TODO Bitbucket Replace pattern
Supported domain list
- github.com
- gitlab.com
 */
func GetRemoteSiteUrl(remoteUrl string) (string, error) {
	// check protocol
	prt := protocol.GetProtocol(remoteUrl)
	if prt == protocol.Unknown {
		return "", &UnknownProtocolError{}
	}

	// switch sequence for domain
	dom := domain.GetDomain(remoteUrl)

	switch dom {
	case domain.Github:
		rep := replaceGithubUrl(remoteUrl)
		return rep, nil
	case domain.GitLab:
		rep := replaceGitlabUrl(remoteUrl)
		return rep, nil
	default:
		return "", &UnknownDomainError{}
	}
}

func replaceGithubUrl(remoteUrl string) string {
	sshForm := "git@github.com:"
	httpsForm := "https://github.com/"
	rep := strings.NewReplacer(sshForm, httpsForm, ".git", "", "\n", "")

	url := rep.Replace(remoteUrl)

	return url
}

func replaceGitlabUrl(remoteUrl string) string {
	sshForm := "git@gitlab.com:"
	httpsForm := "https://gitlab.com/"
	rep := strings.NewReplacer(sshForm, httpsForm, ".git", "", "\n", "")

	url := rep.Replace(remoteUrl)

	return url
}

/*
Methods not currently used
 */
func replaceBitbucketUrl(remoteUrl string) string {
	sshForm := "git@gitlab.com:"
	httpsForm := "https://gitlab.com/"
	rep := strings.NewReplacer(sshForm, httpsForm, ".git", "")

	url := rep.Replace(remoteUrl)

	return url
}
