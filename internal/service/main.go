package service

import "github.com/ebrvkv/adj-parser/internal/model"

type Analyze interface {
	// GetTop10ActiveUsers returns top 10 active users by amount of PRs created and commits pushed
	// (username, PRs count, commits count)
	GetTop10ActiveUsers() []*model.Actor
	// GetTop10RepositoriesCommits returns top 10 repositories by amount of commits pushed (repo name, commits count)
	GetTop10RepositoriesCommits() []*model.Repo
	// GetTop10RepositoriesWatch top 10 repositories by amount of watch events (repo name, watch events count)
	GetTop10RepositoriesWatch() []*model.Repo
}

func NewAnalyzer(
	actors map[uint64]*model.Actor,
	commits map[uint64]*model.CommitStats,
	events map[uint64]*model.Event,
	repos map[uint64]*model.Repo,
) Analyze {
	return &simpleAnalyzer{
		actors:  actors,
		commits: commits,
		events:  events,
		repos:   repos,
	}
}
