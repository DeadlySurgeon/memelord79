package state

import (
	"github.com/go-git/go-git/v5"
)

// Store holds onto information to store between calls.
// Don't forget to save it!
type Store struct {
	repo    *git.Repository
	Example string
	new     bool
	wt      *git.Worktree
}
