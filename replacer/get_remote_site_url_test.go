package replacer_test

import (
	"github.com/cyberfox1002/groo/replacer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRemoteSiteUrl(t *testing.T) {
	// gitlab(SSH): git@gitlab.com:gitlab-org/gitlab-runner.git
	// gitlab(HTTPS): https://gitlab.com/gitlab-org/gitlab-runner.git
	// gitlab(Site): https://gitlab.com/gitlab-org/gitlab-runner
	t.Run("test gitlab service", func(t *testing.T) {
		site := "https://gitlab.com/gitlab-org/gitlab-runner"

		// test ssh
		ssh := "git@gitlab.com:gitlab-org/gitlab-runner.git"
		testSSH, err := replacer.GetRemoteSiteUrl(ssh)
		assert.NoError(t, err)
		assert.Equal(t, site, testSSH)

		// test https
		https := "https://gitlab.com/gitlab-org/gitlab-runner.git"
		testHTTPS, err := replacer.GetRemoteSiteUrl(https)
		assert.NoError(t, err)
		assert.Equal(t, site, testHTTPS)
	})

	// github(SSH): git@github.com:github/training-kit.git
	// github(HTTPS): https://github.com/github/training-kit.git
	// github(Site): https://github.com/github/training-kit
	t.Run("test github service", func(t *testing.T) {
		site := "https://github.com/github/training-kit"

		// test ssh
		ssh := "git@github.com:github/training-kit.git"
		testSSH, err := replacer.GetRemoteSiteUrl(ssh)
		assert.NoError(t, err)
		assert.Equal(t, site, testSSH)

		// test https
		https := "https://github.com/github/training-kit.git"
		testHTTPS, err := replacer.GetRemoteSiteUrl(https)
		assert.NoError(t, err)
		assert.Equal(t, site, testHTTPS)
	})

	// bitbucket(SSH): git@bitbucket.org:bitbucket/bitbucket-theme.git
	// bitbucket(HTTPS): https://username@bitbucket.org/bitbucket/bitbucket-theme.git
	// bitbucket(Site): https://bitbucket.org/bitbucket/bitbucket-theme
	//t.Run("test bitbucket service", func(t *testing.T) {
	//	site := "https://bitbucket.org/bitbucket/bitbucket-theme"
	//
	//	// test ssh
	//	ssh := "git@bitbucket.org:bitbucket/bitbucket-theme.git"
	//	testSSH, err := replacer.GetRemoteSiteUrl(ssh)
	//	assert.NoError(t, err)
	//	assert.Equal(t, site, testSSH)
	//
	//	// test https
	//	https := "https://username@bitbucket.org/bitbucket/bitbucket-theme.git"
	//	testHTTPS, err := replacer.GetRemoteSiteUrl(https)
	//	test := strings.TrimPrefix(testHTTPS, "username")
	//	assert.NoError(t, err)
	//	assert.Equal(t, site, test)
	//})
}
