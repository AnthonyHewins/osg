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

	"io"
	"io/ioutil"
	"path/filepath"

	"archive/zip"
	"archive/tar"
	"compress/gzip"
)

type Crawlable = func(*string, chan Option) error

func crawl_dir(path *string, files chan Option) error {
	return filepath.Walk(*path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil || whitelisted(info) { return err }

			file_ptr, err := os.Open(path)
			if err != nil { return err }

			buf, err := ioutil.ReadAll(file_ptr)
			if err != nil { return err }

			f := File{path, buf}
			files <- Option{File: &f, Err: nil}
			return nil
		},
	)
}

func crawl_zip(filename *string, files chan Option) error {
	r, err := zip.OpenReader(*filename)
	if err != nil { return err }

	defer r.Close()

	for _, reader := range r.File {
		if os_fileinfo := reader.FileHeader.FileInfo(); whitelisted(os_fileinfo) {
			continue
		}

		current_file, err := reader.Open()
		if err != nil { return err }

		buf, err := ioutil.ReadAll(current_file)
		if err != nil { return err }

		f := File{reader.FileHeader.Name, buf}
		files <- Option{File: &f, Err: err}
	}

	return nil
}

func crawl_tar(filename *string, files chan Option) error {
	f, err := os.Open(*filename)
	if err != nil { return err }

	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil { return err }

	defer gzf.Close()

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		switch err {
		case io.EOF:
			return nil
		case nil:
			// no-op, this case means nothing's wrong
		default:
			return err
		}

		// PAX headers (XHeaders shown below) are just file permissions, which we don't care about
		this_is_case_we_dont_care_about := header.Typeflag == tar.TypeXHeader ||
			header.Typeflag == tar.TypeXGlobalHeader ||
			whitelisted(header.FileInfo())

		if this_is_case_we_dont_care_about {
			continue
		}

		data := make([]byte, header.Size)
		if _, err := tarReader.Read(data); err == io.EOF {
			f := File{ header.Name, data }
			files <- Option{File: &f, Err: err}
		} else if err != nil {
			return err
		}
	}
}
