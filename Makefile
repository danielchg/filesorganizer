
build:
	go build .

all: clean build test

.PHONY=test
test:
	./filesorganizer --src ./samples --dest ./test

.PHONY=clean
clean:
	rm -rf ./test
	rm -f filesorganizer
