package fetch

import (
	"log"
)

// TODO add github capabilities and the API
func StartDataPipeline(uri *string, pipeline chan File) {
	location, err := where_to_go(uri)
	if err != nil { log.Fatalln(err) }

	switch location {
	case LOCAL_DIRECTORY:
		crawl_dir(uri, pipeline)
	case LOCAL_TAR:
		inflate_tar(*uri, pipeline)
	case LOCAL_ZIP:
		inflate_zip(*uri, pipeline)
	case REMOTE_ZIP:
		filename, err := download_to_temp(uri)
		if err != nil { log.Fatalln(err) }

		inflate_zip(filename, pipeline)
	case REMOTE_TAR:
		filename, err := download_to_temp(uri)
		if err != nil { log.Fatalln(err) }

		inflate_tar(filename, pipeline)
	}
}
