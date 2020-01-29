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

type File struct {
	Name string
	Contents []byte
}

func crawl_dir(path *string, pipeline chan File) {
	err := filepath.Walk(*path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil { return err }

			if info.IsDir() { return nil }

			filename := info.Name()
			if blacklisted(&filename) { return nil }

			absolute_path, err := filepath.Abs(path)
			if err != nil { return err }

			file_ptr, err := os.Open(absolute_path)
			if err != nil { return err }

			buf, err := ioutil.ReadAll(file_ptr)
			if err != nil { return err }

			pipeline <- File{filename, buf}
			return nil
		},
	)

	if err != nil { log.Fatalln(err) }

	close(pipeline)
}

//
func inflate_zip(filename string, pipeline chan File) {
	r, err := zip.OpenReader(filename)
	if err != nil { log.Fatalln(err) }

	defer r.Close()

	for _, reader := range r.File {
		if filter_zip(&reader.FileHeader) {
			current_file, err := reader.Open()

			if err != nil { log.Fatalf("couldn't open (filename is '%v'?): %v", current_file, err) }

			buf, err := ioutil.ReadAll(current_file)
			if err != nil { log.Fatalf("ran into error reading file %v: %v", current_file, err) }

			todo := File{reader.FileHeader.Name, buf}
			pipeline <- todo
		}
	}

	close(pipeline)
}

func inflate_tar(filename string, pipeline chan File) {
	f, err := os.Open(filename)
	if err != nil { log.Fatalln(err) }

	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil { log.Fatalln(err) }

	tarReader := tar.NewReader(gzf)

	for {
		header, err := tarReader.Next()

		if err == io.EOF { break            }
		if err != nil    { log.Fatalln(err) }

		if filter_tar(header) {

			data := make([]byte, header.Size)
			_, err := tarReader.Read(data)

			switch err {
			case io.EOF:
				pipeline <- File{ header.Name, data }
			case nil:
				continue
			default:
				log.Fatalln(err)
			}
		}
	}

	close(pipeline)
}
