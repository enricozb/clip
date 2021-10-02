package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"

	"github.com/enricozb/clip/lib/clip"
)

func main() {
	var content []byte
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		content, _ = io.ReadAll(os.Stdin)
	} else {
		content = clip.Get()
	}

	if _, err := os.Stdout.Write(content); err != nil {
		panic(fmt.Errorf("stdout write: %v", err))
	}
	if err := os.Stdout.Close(); err != nil {
		panic(fmt.Errorf("stdout close: %v", err))
	}

	clip.Sync(content)
}
