GO_FILES := $(shell find -type f -name '*.go')

.PHONY: clean

all: clip daemon

clip: $(GO_FILES)
	go build -o ./bin/clip cmd/clip/main.go

daemon: $(GO_FILES)
	go build -o ./bin/clipd cmd/daemon/main.go

clean:
	rm -r bin
