package fetch

import (
	"testing"
	"io/ioutil"
	"path/filepath"
)

func TestDownloadToTmp(t *testing.T) {
	// ioutil.TempFile("", string) shows us where tmp files live ("" means put it in the OS tmp dir)
	tmpfile, err := ioutil.TempFile("", "tmpfile-test-*")
	if err != nil {
		t.Errorf("unable to create a tmpfile for tests: %v", err)
		return
	}

	example_repo := "https://github.com/AnthonyHewins/one-time-pad-socket/archive/master.zip"

	filename, err := download_to_tmp(&example_repo)
	if err != nil {
		t.Errorf("Unable to download, check inet related issues first: %v", err)
		return
	}

	if filepath.Dir(filename) != filepath.Dir(tmpfile.Name()) {
		t.Errorf("the file isn't in the temp dir of the OS (actual != expected) : %v != %v", filename, tmpfile)
	}
}
