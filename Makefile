build:
	go build main.go filereader.go jsonhandler.go manager.go worker.go

configure:
	go get github.com/alexcesaro/log

clean:
	rm -rf main
