package clip

import (
	"github.com/enricozb/clip/lib/config"
)

// Get returns the first successful clipboard retrieval from the configured clipboards.
func Get() []byte {
	for _, clipboard := range config.Clipboards {
		if content, err := clipboard.Get(); err == nil {
			return content
		}
	}

	return nil
}

// Sync sets the clipboard contents of all configured clipboards
func Sync(content []byte) {
	for _, clipboard := range config.Clipboards {
		clipboard.Set(content)
	}
}
