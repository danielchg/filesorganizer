# build the binary
build:
	go mod download
	go build .

# run an example for testing
.PHONY: run
run:
	./filesorganizer --src ./samples --dest ./test

# clean all the files created by a test run
.PHONY: clean
clean:
	rm -rf ./test
	rm -f filesorganizer

# run the command using the binary
all: clean build run

# run unit tests
.PHONY: test
test:
	go test ./... --cover
