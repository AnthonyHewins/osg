// This file exposes several functions that do one of the following:
//
//   1. Inflate compressed media that was downloaded or is stored locally in-place,
//      buffering one file at a time in memory, and sending it down the pipe
//   2. Scanning a dir for files, reading them to the end, then sending those down the pipe
//
// The files are prepared using the struct below for later processing. Note some are filtered out.

package fetch

import (
	"os"
	"log"

	"io"
	"io/ioutil"
	"path/filepath"

	"archive/zip"
	"archive/tar"
	"compress/gzip"
)

func crawl_dir(path *string, files chan File, errors chan error) {
	filepath.Walk(*path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				errors <- err
				return nil
			}

			if whitelisted(info) { return nil }

			file_ptr, err := os.Open(path)
			if err != nil {
				errors <- err
				return nil
			}

			buf, err := ioutil.ReadAll(file_ptr)
			if err != nil {
				errors <- err
				return nil
			}

			files <- File{info.Name(), buf}
			return nil
		},
	)
}

func crawl_zip(filename *string, files chan File, errors chan error) {
	r, err := zip.OpenReader(*filename)
	if err != nil {
		log.Fatalln(err)
	}

	defer r.Close()

	for _, reader := range r.File {
		if os_fileinfo := reader.FileHeader.FileInfo(); whitelisted(os_fileinfo) {
			continue
		}

		current_file, err := reader.Open()
		if err != nil {
			errors <- err
			continue
		}

		buf, err := ioutil.ReadAll(current_file)
		if err != nil {
			errors <- err
			continue
		}

		files <- File{reader.FileHeader.Name, buf}
	}
}

func crawl_tar(filename *string, files chan File, errors chan error) {
	f, err := os.Open(*filename)
	if err != nil {
		errors <- err
		return
	}

	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		errors <- err
		return
	}

	defer gzf.Close()

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		switch err {
		case io.EOF:
			return
		case nil:
			// no-op, this case means nothing's wrong
		default:
			errors <- err
			continue
		}

		if whitelisted(header.FileInfo()) {
			continue
		}

		data := make([]byte, header.Size)
		if _, err := tarReader.Read(data); err == io.EOF {
			files <- File{ header.Name, data }
		} else if err != nil {
			errors <- err
		}
	}
}
