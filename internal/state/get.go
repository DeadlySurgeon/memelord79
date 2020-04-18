package state

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

// Get retrieves the state from the repository
func Get(repoURL string) (*Store, error) {
	state := &Store{}

	// Clone/Open Repository
	repo, err := getRepository(repoURL, auth)
	state.repo = repo

	defaultRef, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("Failed to get HEAD: %w", err)
	}

	ref := plumbing.NewHashReference("refs/heads/state", defaultRef.Hash())

	if err = repo.Storer.SetReference(ref); err != nil {
		return nil, err
	}

	if err = createRemoteStateBranch(repo, auth); err != nil {
		return nil, fmt.Errorf("Failed to create branch on remote: %w", err)
	}

	if state.wt, err = repo.Worktree(); err != nil {
		return nil, fmt.Errorf("Failed to aquire a worktree: %w", err)
	}

	if err = state.wt.Checkout(&git.CheckoutOptions{
		Branch: ref.Name(),
	}); err != nil {
		return nil, fmt.Errorf("Failed to checkout state: %w", err)
	}

	stateReader, err := os.Open("./state/state.json")
	if err != nil {
		if os.IsNotExist(err) {
			return state, nil
		}
		return nil, err
	}
	defer stateReader.Close()

	return state, json.NewDecoder(stateReader).Decode(state)
}

func createRemoteStateBranch(repo *git.Repository, auth transport.AuthMethod) error {
	refs, err := repo.References()
	if err != nil {
		return fmt.Errorf("Failed to check refs: %w", err)
	}
	exist := false
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().String() == "refs/remotes/origin/state" {
			exist = true
		}
		return nil
	})
	if !exist {
		if err = repo.Push(&git.PushOptions{
			Auth: auth,
			// RefSpecs: []config.RefSpec{
			// 	"+refs/heads/state:refs/remotes/origin/state",
			// },
		},
		); err != nil {
			return fmt.Errorf("Failed to push remote repo: %w", err)
		}
	}

	return nil
}
