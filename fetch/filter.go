package fetch

import (
	"os"
	"path/filepath"
)

// TODO this isnt even close to scalable, create a true black/blacklist
// loaded from external source or something
var blacklist_files = map[string]bool{
	".gitignore": true,
	".gitattributes": true,
	"Cargo.toml": true,
	"Cargo.lock": true,
}

var blacklist_extensions = map[string]bool {
	".png":  true,
	".mp3":  true,
	".jpg":  true,
	".jpeg": true,
	".md":   true, // readmes are typically not going to be filled with deceptive code
}

func whitelisted(metadata os.FileInfo) bool {
	if metadata.IsDir() { return true }

	filename := metadata.Name()
	return blacklist_files[filename] ||
		blacklist_extensions[filepath.Ext(filename)]
}
