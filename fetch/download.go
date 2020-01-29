package fetch

import (
	"net/http"
	"io"
	"io/ioutil"
)

const tmpfile_name = "osg-repo-*"

func download_to_tmp(uri *string) (string, error) {
	resp, err := http.Get(*uri)
	if err != nil { return tmpfile_name, err }

	f, err := ioutil.TempFile("", tmpfile_name)
	if err != nil { return tmpfile_name, err }

	_, err = io.Copy(f, resp.Body)

	resp.Body.Close()
	f.Close()

	return f.Name(), err
}
