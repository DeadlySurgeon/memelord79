package state

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5/plumbing/transport"
	gitSSH "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

// Don't use singletons. If state is moved outside of internal, this has to go.
var auth transport.AuthMethod

func init() {
	// try to load a default key to use
	auth, _ = Authenticate(nil)
}

// Authenticate gets the key and prepares SSH for go-git to use.
func Authenticate(key []byte) (*gitSSH.PublicKeys, error) {
	if key == nil {
		keySTR := os.Getenv("MEMELORD_SSH_PRIVATE_KEY")
		if keySTR == "" {
			return nil, fmt.Errorf("ENV `MEMELORD_SSH_PRIVATE_KEY` is not set with a private RSA key")
		}
		key = []byte(keySTR)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	auth := &gitSSH.PublicKeys{
		Signer: signer,
		HostKeyCallbackHelper: gitSSH.HostKeyCallbackHelper{
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
		User: "git",
	}
	return auth, nil
}
