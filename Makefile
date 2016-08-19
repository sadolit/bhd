

.PHONY : all
all : build test

.PHONY : test
test :
	go test

.PHONY : build
build: bindir
	go build -o ./bin/hbd

.PHONY: bindir
bindir:
	@mkdir -p bin
