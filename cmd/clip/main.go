package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"

	"github.com/enricozb/clip/lib/config"
)

func getClipboardContent() []byte {
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		content, _ := io.ReadAll(os.Stdin)
		return content
	}

	for _, clipboard := range config.Clipboards {
		if content, err := clipboard.Get(); err == nil {
			return content
		}
	}

	return nil
}

func syncClipboardContent(content []byte) {
	// update clipboards
	for _, clipboard := range config.Clipboards {
		clipboard.Set(content)
	}
}

func main() {
	content := getClipboardContent()

	if _, err := os.Stdout.Write(content); err != nil {
		panic(fmt.Errorf("stdout write: %v", err))
	}
	if err := os.Stdout.Close(); err != nil {
		panic(fmt.Errorf("stdout close: %v", err))
	}

	syncClipboardContent(content)
}
