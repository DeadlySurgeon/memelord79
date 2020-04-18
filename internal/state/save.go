package state

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Save puts the state back into the repository
func Save(repo string, store *Store) error {
	stateWriter, err := os.Open("./state/state.json")
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("Failed to open state: %w", err)
		}
		if stateWriter, err = os.Create("./state/state.json"); err != nil {
			return fmt.Errorf("Failed to create state file: %w", err)
		}
	}
	defer stateWriter.Close()

	if err = json.NewEncoder(stateWriter).Encode(store); err != nil {
		return fmt.Errorf("Failed to encode state: %w", err)
	}

	// TODO:
	// - Delete history. It will just make the dumb repo bigger than it needs to be.
	// - check if it was modified before last add.
	//   + Causes an issue otherwise.
	if _, err = store.wt.Add("state.json"); err != nil {
		return fmt.Errorf("Couldn't add state to tracking: %w", err)
	}

	commit, err := store.wt.Commit("Updated state", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "The Surgeon",
			Email: "deadly.surgery@gmail.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return fmt.Errorf("Failed to generate commit: %w", err)
	}

	_, err = store.repo.CommitObject(commit)
	if err != nil {
		return fmt.Errorf("Failed to commit work: %w", err)
	}

	if err = store.repo.Push(&git.PushOptions{
		Progress: os.Stdout,
		RefSpecs: []config.RefSpec{
			"+refs/heads/state:refs/heads/state",
		},
		Auth: auth,
	}); err != nil {
		return fmt.Errorf("Failed to push repo: %w", err)
	}

	return nil
}
