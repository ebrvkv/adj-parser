package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ebrvkv/adj-parser/internal/controller"

	"github.com/ebrvkv/adj-parser/internal/service"
	"github.com/golang/glog"

	"github.com/ebrvkv/adj-parser/internal/repository"
)

// Version is a compile time variable for CI/CD pipeline
var Version string

func main() {
	actorsPath := flag.String("a", "assets/data/actors.csv", "path to csv file with actors")
	commitsPath := flag.String("c", "assets/data/commits.csv", "path to csv file with commits")
	eventsPath := flag.String("e", "assets/data/events.csv", "path to csv file with events")
	reposPath := flag.String("r", "assets/data/repos.csv", "path to csv file with repositories")
	flag.Parse()

	repo := repository.NewFilesRepo(*actorsPath, *commitsPath, *eventsPath, *reposPath)

	actors, err := repo.GetActors()
	if err != nil {
		glog.Fatal(err)
	}

	events, err := repo.GetEvents()
	if err != nil {
		glog.Fatal(err)
	}

	repos, err := repo.GetRepos()
	if err != nil {
		log.Fatal(err)
	}

	commits, err := repo.GetCommits()
	if err != nil {
		glog.Fatal(err)
	}

	analyzer := service.NewAnalyzer(actors, commits, events, repos)

	fmt.Println(controller.PrettifyActors(analyzer.GetTop10ActiveUsers()))
	fmt.Println(controller.PrettifyReposCommits(analyzer.GetTop10RepositoriesCommits()))
	fmt.Println(controller.PrettifyReposWatchEvents(analyzer.GetTop10RepositoriesWatch()))
}
