package auth

import (
	"os"
	"path"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var (
	testPrivateKey     = []byte("very_sekrit_key")
	testToken          = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.AqFWnFeY9B8jj7-l3z0a9iaZdwIca7xhUF3fuaJjU90"
	testInertiaKeyPath = path.Join(os.Getenv("GOPATH"), "/src/github.com/ubclaunchpad/inertia/test/keys/id_rsa")
)

// Helper function that implements jwt.keyFunc
func getFakeAPIKey(tok *jwt.Token) (interface{}, error) {
	return testPrivateKey, nil
}

func TestGetAPIPrivateKey(t *testing.T) {
	key, err := getAPIPrivateKeyFromPath(nil, testInertiaKeyPath)
	assert.Nil(t, err)
	assert.Contains(t, string(key.([]byte)), "user: git, name: ssh-public-keys")
}

func TestGetGithubKey(t *testing.T) {
	pemFile, err := os.Open(testInertiaKeyPath)
	assert.Nil(t, err)
	_, err = GetGithubKey(pemFile)
	assert.Nil(t, err)
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(testPrivateKey)
	assert.Nil(t, err, "generateToken must not fail")
	assert.Equal(t, token, testToken)

	otherToken, err := GenerateToken([]byte("another_sekrit_key"))
	assert.Nil(t, err)
	assert.NotEqual(t, token, otherToken)
}
