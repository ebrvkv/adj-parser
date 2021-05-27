package repository

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang/glog"

	"github.com/ebrvkv/adj-parser/internal/model"
)

type filesRepo struct {
	actorsPath  string
	commitsPath string
	eventsPath  string
	repoPath    string
}

func newFileRepo(actorsPath, commitsPath, eventsPath, repoPath string) *filesRepo {
	return &filesRepo{
		actorsPath:  actorsPath,
		commitsPath: commitsPath,
		eventsPath:  eventsPath,
		repoPath:    repoPath,
	}
}

func readCSVRows(path string) (*csv.Reader, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	return csv.NewReader(f), nil
}

func (fr *filesRepo) GetActors() (map[uint64]*model.Actor, error) {
	rows, err := readCSVRows(fr.actorsPath)
	if err != nil {
		return nil, err
	}
	res := make(map[uint64]*model.Actor)
	var flag bool
	for {
		// Ignore first line
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if len(row) < 2 {
			glog.Error("Bad line")
		}
		if !flag {
			flag = true
			continue
		}

		id, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		res[id] = &model.Actor{
			ID:       id,
			Username: row[1],
		}
	}
	return res, nil
}

func (fr *filesRepo) GetCommits() (map[uint64]*model.CommitStats, error) {
	res := make(map[uint64]*model.CommitStats)
	dedup := make(map[string]struct{})
	rows, err := readCSVRows(fr.commitsPath)
	if err != nil {
		return nil, err
	}
	var flag bool
	var cnt uint
	for {
		cnt++
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if len(row) < 3 {
			glog.Error("Bad line")
		}
		if !flag {
			flag = true
			continue
		}
		eID, err := strconv.ParseUint(row[2], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		if _, ok := dedup[row[0]]; ok {
			continue
		}
		dedup[row[0]] = struct{}{}
		if _, ok := res[eID]; !ok {
			res[eID] = &model.CommitStats{}
		}
		res[eID].Commits++
		if strings.Contains(row[1], "Merge pull request #") {
			res[eID].PRs++
		}
	}
	return res, err
}

func (fr *filesRepo) GetEvents() (map[uint64]*model.Event, error) {
	res := make(map[uint64]*model.Event)
	rows, err := readCSVRows(fr.eventsPath)
	if err != nil {
		return nil, err
	}
	var flag bool
	for {
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if len(row) < 4 {
			glog.Error("Bad line")
		}
		if !flag {
			flag = true
			continue
		}
		eventID, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		actorID, err := strconv.ParseUint(row[2], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		repoID, err := strconv.ParseUint(row[3], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		res[eventID] = &model.Event{
			ActorID: actorID,
			RepoID:  repoID,
			Type:    row[1],
		}
	}
	return res, err
}

func (fr *filesRepo) GetRepos() (map[uint64]*model.Repo, error) {
	res := make(map[uint64]*model.Repo)
	rows, err := readCSVRows(fr.repoPath)
	if err != nil {
		return nil, err
	}
	var flag bool
	for {
		row, err := rows.Read()
		if err == io.EOF {
			break
		}
		if len(row) < 2 {
			glog.Error("Bad line")
		}
		if !flag {
			flag = true
			continue
		}
		id, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			glog.Error(err)
			continue
		}
		res[id] = &model.Repo{
			ID:   id,
			Name: row[1],
		}
	}
	return res, err
}
