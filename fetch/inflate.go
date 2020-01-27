package fetch

import (
	"os"
	"io/ioutil"
	"io"
	"log"
	"archive/zip"
	"archive/tar"
	"compress/gzip"
)

type File struct {
	Name string
	Contents []byte
}

func inflate_zip(filename string, pipeline chan File) {
	r, err := zip.OpenReader(filename)
	if err != nil { log.Fatalf("ran into zip error: %v\n", err) }

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
