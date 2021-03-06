package common

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var remoteURLVariations = []string{
	// SSH URL
	"git@github.com:ubclaunchpad/inertia.git",
	"git@gitlab.com:ubclaunchpad/inertia.git",
	"git@bitbucket.org:ubclaunchpad/inertia.git",

	// Github URL Variations
	"https://github.com/ubclaunchpad/inertia.git",
	"git://github.com/ubclaunchpad/inertia.git",

	// Gitlab URL Variations
	"https://gitlab.com/ubclaunchpad/inertia.git",
	"git://gitlab.com/ubclaunchpad/inertia.git",

	// Bitbucket URL Variations
	"https://ubclaunchpad@bitbucket.org/ubclaunchpad/inertia.git",
}

func TestCheckForGit(t *testing.T) {
	cwd, err := os.Getwd()
	assert.Nil(t, err)
	assert.NotEqual(t, nil, CheckForGit(cwd))
	inertia, _ := path.Split(cwd)
	assert.Equal(t, nil, CheckForGit(inertia))
}

func TestGetSSHRemoteURL(t *testing.T) {
	validSSH := remoteURLVariations[0:3]
	for _, url := range remoteURLVariations {
		assert.Contains(t, validSSH, GetSSHRemoteURL(url))
	}
}

func TestGetBranchFromRef(t *testing.T) {
	branch := GetBranchFromRef("refs/heads/master")
	assert.Equal(t, "master", branch)
}
