package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	gHTTP "github.com/go-git/go-git/v5/plumbing/transport/http"
	gitSSH "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

func main() {

	repoURL := "git@github.com:DeadlySurgeon/memerlord79.git"

	repo, err := git.PlainClone("/store",
		false,
		&git.CloneOptions{
			Auth: &gHTTP.BasicAuth{
				Username: "memerlord79",
				Password: "<token>",
			},
			URL:      repoURL,
			Progress: os.Stdout,
		},
	)
	if err != nil {
		panic(err)
	}

	_ = repo

	fact, err := getFact()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get a printer fact: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(fact)
}

func getFact() (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://printerfacts.cetacean.club/fact",
		nil,
	)
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	factBytes, err := ioutil.ReadAll(resp.Body)
	return string(factBytes), err
}

var repo = "git@github.com:DeadlySurgeon/memerlord79.git"

// GetState returns a state from the repo
func GetState(repo string) (*State, error) {
	auth, err := getAuth()
	if err != nil {
		return nil, err
	}
	gitRepo, err := git.PlainClone("./state", false, &git.CloneOptions{
		Auth:          auth,
		URL:           repo,
		ReferenceName: "state",
	})
	if err != nil {
		return nil, err
	}

	var state *State

	return state, nil
}

// SaveState stores the state back to the repo
func SaveState(repo string, state *State) error {

	return nil
}

// State holds onto information to store between calls.
// Don't forget to save it!
type State struct {
	repo *git.Repository
}

func getAuth() (*gitSSH.PublicKeys, error) {
	key := os.Getenv("SSH_PRIVATE_KEY")
	if key == "" {
		return nil, fmt.Errorf("ENV `SSH_PRIVATE_KEY` is not set with a private RSA key")
	}

	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}

	auth := &gitSSH.PublicKeys{
		Signer: signer,
		User:   "deadly.surgery@gmail.com",
	}
	return auth, nil
}
