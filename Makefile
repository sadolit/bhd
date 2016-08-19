

.PHONY : all
all : build test

.PHONY : test
test :
	go test

.PHONY : build
build: bindir
	go build -o ./bin/bhd

.PHONY: bindir
bindir:
	@mkdir -p bin

.PHONY: clean
clean:
	@rm -rf bin
