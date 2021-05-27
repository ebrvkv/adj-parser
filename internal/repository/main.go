package repository

import "github.com/ebrvkv/adj-parser/internal/model"

// GitHub is the main interface for stats
type GitHub interface {
	// GetActors returns a list of actors from the repo
	GetActors() (map[uint64]*model.Actor, error)
	// GetCommits returns a list of commits from the repo
	GetCommits() (map[uint64]*model.CommitStats, error)
	// GetEvents returns a list of events from the repo
	GetEvents() (map[uint64]*model.Event, error)
	// GetRepos returns a list of repositories from the repo
	GetRepos() (map[uint64]*model.Repo, error)
}

// NewFilesRepo returns repository with data gathered from the files
func NewFilesRepo(actorsPath, commitsPath, eventsPath, repoPath string) GitHub {
	return newFileRepo(actorsPath, commitsPath, eventsPath, repoPath)
}
