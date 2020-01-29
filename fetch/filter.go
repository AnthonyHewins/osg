package fetch

import (
	"archive/zip"
	"archive/tar"
	"path/filepath"
)

// TODO this isnt even close to scalable, create a true black/blacklist
// loaded from external source or something
var blacklist_files = map[string]bool{
	".gitignore": true,
	".gitattributes": true,
}

var blacklist_extensions = map[string]bool {
	".png":  true,
	".mp3":  true,
	".jpg":  true,
	".jpeg": true,
	".md":   true, // readmes are typically not going to be filled with deception
}

func filter_zip(header *zip.FileHeader) bool {
	return !blacklisted(&header.Name)
}

func filter_tar(header *tar.Header) bool {
	if header.Typeflag != tar.TypeReg { return false }
	return !blacklisted(&header.Name)
}

func blacklisted(file *string) bool {
	filename := filepath.Base(*file)

	return blacklist_files[filename] ||
		blacklist_extensions[filepath.Ext(filename)]
}
