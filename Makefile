
build:
	go build .



.PHONY=run
run:
	./filesorganizer --src ./samples --dest ./test

.PHONY=clean
clean:
	rm -rf ./test
	rm -f filesorganizer

all: clean build run

.PHONY=test
test:
	go test ./... --cover
