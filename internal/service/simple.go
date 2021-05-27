package service

import (
	"sort"

	"github.com/ebrvkv/adj-parser/internal/model"
)

type simpleAnalyzer struct {
	actors  map[uint64]*model.Actor
	commits map[uint64]*model.CommitStats
	events  map[uint64]*model.Event
	repos   map[uint64]*model.Repo
}

type ActorStats struct {
	Username string
	PRs      uint64
	Commits  uint64
}

func (s *simpleAnalyzer) GetTop10ActiveUsers() (res []*model.Actor) {
	for eID := range s.commits {
		s.actors[s.events[eID].ActorID].Commits += s.commits[eID].Commits
		s.actors[s.events[eID].ActorID].PRs += s.commits[eID].PRs
	}
	//for eID, e := range s.events {
	//	if e.Type == "PullRequestEvent" {
	//		s.actors[s.events[eID].ActorID].PRs++
	//	}
	//}
	for _, r := range s.actors {
		res = append(res, r)
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i].PRs == res[j].PRs {
			return res[i].Commits > res[j].Commits
		}
		return res[i].PRs > res[j].PRs
	})
	if len(res) > 10 {
		return res[:10]
	}
	return res
}

func (s *simpleAnalyzer) GetTop10RepositoriesCommits() (res []*model.Repo) {
	for eID := range s.commits {
		s.repos[s.events[eID].RepoID].Commits += s.commits[eID].Commits
	}
	for _, r := range s.repos {
		res = append(res, r)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Commits > res[j].Commits
	})
	if len(res) > 10 {
		return res[:10]
	}
	return res
}

func (s *simpleAnalyzer) GetTop10RepositoriesWatch() (res []*model.Repo) {
	for eID, e := range s.events {
		if e.Type == "WatchEvent" {
			s.repos[s.events[eID].RepoID].WatchEvents++
		}
	}
	for _, r := range s.repos {
		res = append(res, r)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].WatchEvents > res[j].WatchEvents
	})
	if len(res) > 10 {
		return res[:10]
	}
	return res
}
