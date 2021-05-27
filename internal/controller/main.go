package controller

import (
	"fmt"

	"github.com/ebrvkv/adj-parser/internal/model"
)

func PrettifyReposCommits(repos []*model.Repo) (res string) {
	res = "Top 10 repositories by amount of commits pushed:\n\n"
	for i, repo := range repos {
		res += fmt.Sprintf("%d. %s - %d commits\n", i+1, repo.Name, repo.Commits)
	}
	return res
}

func PrettifyReposWatchEvents(repos []*model.Repo) (res string) {
	res = "Top 10 repositories by amount of watch events:\n\n"
	for i, repo := range repos {
		res += fmt.Sprintf("%d. %s - %d watch events\n", i+1, repo.Name, repo.WatchEvents)
	}
	return res
}

func PrettifyActors(actors []*model.Actor) (res string) {
	res = "Top 10 active users by amount of PRs created and commits pushed:\n\n"
	for i, actor := range actors {
		res += fmt.Sprintf("%d. %s - %d PRs, %d commits\n", i+1, actor.Username, actor.PRs, actor.Commits)
	}
	return res
}
