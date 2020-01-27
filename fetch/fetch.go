package fetch

import (
	"log"
)

type DownloadError struct {
	msg string
}

func (err DownloadError) Error() string {
	return err.msg
}

type CompressionType uint8

const (
	Tarball CompressionType = 0
	Zipball CompressionType = 1
)

// TODO add github capabilities and the API
func StartDataPipeline(uri *string, pipeline chan File) {
	extension, err := get_extension(uri)
	if err != nil { log.Fatalln(err) }

	filename, err := download(uri, extension)
	if err != nil { log.Fatalln(err) }

	if extension == Tarball {
		inflate_tar(filename, pipeline)
	} else {
		inflate_zip(filename, pipeline)
	}
}

func get_extension(uri *string) (CompressionType, error) {
	n := len(*uri)
	if (*uri)[n-7:] == ".tar.gz" { return Tarball, nil }
	if (*uri)[n-4:] == ".zip"    { return Zipball, nil }

	return Tarball, DownloadError{"unable to find file extension"}
}
