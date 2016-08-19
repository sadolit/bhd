

.PHONY : all
all : bindir
	go build -o ./bin/hbd

.PHONY: bindir
bindir:
	@mkdir -p bin
