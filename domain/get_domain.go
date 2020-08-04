package domain

import "strings"

type Domain int

const (
	Github Domain = iota
	GitLab
	BitBucket
	Unknown
)

const (
	githubDomain = "github.com"
	gitlabDomain = "gitlab.com"
	bitbucketDomain = "bitbucket.org"
)

func GetDomain(remoteUrl string) Domain {
	if strings.Contains(remoteUrl, githubDomain) {
		return Github
	}

	if strings.Contains(remoteUrl, gitlabDomain) {
		return GitLab
	}

	if strings.Contains(remoteUrl, bitbucketDomain) {
		return BitBucket
	}

	return Unknown
}
