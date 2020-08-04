package domain_test

import (
	"github.com/cyberfox1002/groo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDomain(t *testing.T) {
	var dom domain.Domain

	// gitlab(SSH): git@gitlab.com:gitlab-org/gitlab-runner.git
	// gitlab(HTTPS): https://gitlab.com/gitlab-org/gitlab-runner.git
	// gitlab(Site): https://gitlab.com/gitlab-org/gitlab-runner
	t.Run("test gitlab service", func(t *testing.T) {
		// test ssh
		ssh := "git@gitlab.com:gitlab-org/gitlab-runner.git"
		dom = domain.GetDomain(ssh)
		assert.Equal(t, dom, domain.GitLab)

		// test https
		https := "https://gitlab.com/gitlab-org/gitlab-runner.git"
		dom = domain.GetDomain(https)
		assert.Equal(t, dom, domain.GitLab)
	})

	// github(SSH): git@github.com:github/training-kit.git
	// github(HTTPS): https://github.com/github/training-kit.git
	// github(Site): https://github.com/github/training-kit
	t.Run("test github service", func(t *testing.T) {
		// test ssh
		ssh := "git@github.com:github/training-kit.git"
		dom = domain.GetDomain(ssh)
		assert.Equal(t, dom, domain.Github)

		// test https
		https := "https://github.com/github/training-kit.git"
		dom = domain.GetDomain(https)
		assert.Equal(t, dom, domain.Github)
	})

	// bitbucket(SSH): git@bitbucket.org:bitbucket/bitbucket-theme.git
	// bitbucket(HTTPS): https://username@bitbucket.org/bitbucket/bitbucket-theme.git
	// bitbucket(Site): https://bitbucket.org/bitbucket/bitbucket-theme
	t.Run("test bitbucket service", func(t *testing.T) {
		// test ssh
		ssh := "git@bitbucket.org:bitbucket/bitbucket-theme.git"
		dom = domain.GetDomain(ssh)
		assert.Equal(t, dom, domain.BitBucket)

		// test https
		https := "https://username@bitbucket.org/bitbucket/bitbucket-theme.git"
		dom = domain.GetDomain(https)
		assert.Equal(t, dom, domain.BitBucket)
	})
}
