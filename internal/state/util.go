package state

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
)

func getRepository(repoURL string, auth transport.AuthMethod) (repo *git.Repository, err error) {
	// Clone/Open Repository
	repo, err = git.PlainClone("./state", false, &git.CloneOptions{
		Auth: auth,
		URL:  repoURL,
	})
	if err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return nil, fmt.Errorf("Failed to clone repo: %w", err)
		}
		if repo, err = git.PlainOpen("state"); err != nil {
			return nil, fmt.Errorf("Failed to open existing repo: %w", err)
		}
	}
	return repo, err
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
