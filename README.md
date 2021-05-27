# adj-parser

## Goal

The goal is to write a script that outputs:

- Top 10 active users by amount of PRs created and commits pushed (username, PRs count, commits count)
- Top 10 repositories by amount of commits pushed (repo name, commits count)
- Top 10 repositories by amount of watch events (repo name, watch events count)

This assignment must be done in any type-safe language, that the candidate prefers.

## How to run

From the root directory of the repo run `go run cmd/parser/main.go`. By default, all csv files from `assets/data` dir
will be used.

### CLI arguments

```
-a string
    path to csv file with actors (default "assets/data/actors.csv")

-c string
    path to csv file with commits (default "assets/data/commits.csv")

-e string
    path to csv file with events (default "assets/data/events.csv")

-r string
    path to csv file with repositories (default "assets/data/repos.csv")
```



