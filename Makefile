# https://www.youtube.com/watch?v=QztvWSCbQLU

hello:
	echo "hello"

build:
	go build -o bin/main main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

all: compile hello