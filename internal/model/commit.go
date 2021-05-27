package model

type Commit struct {
	Sha     string
	Message string
	EventID uint64
}

type CommitStats struct {
	PRs     uint64
	Commits uint64
}
