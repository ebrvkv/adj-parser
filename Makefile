run-defaults:
	go run cmd/parser/main.go -a assets/data/actors.csv -c assets/data/commits.csv -e assets/data/events.csv -r assets/data/repos.csv

build:
	go build -o github-events cmd/parser/main.go
