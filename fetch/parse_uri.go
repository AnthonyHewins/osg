package fetch

import (
	"fmt"
	"os"
	"path"
)

type WhereToFetch = uint8

const (
	LOCAL_DIRECTORY = iota
	LOCAL_TAR       = iota
	LOCAL_ZIP       = iota
	REMOTE_ZIP      = iota
	REMOTE_TAR      = iota
)

type PathError struct {
	msg string
}

func (p PathError) Error() string {
	return p.msg
}

func where_to_go(uri *string) (WhereToFetch, error) {
	if fileinfo, err := os.Stat(*uri); os.IsNotExist(err) {
		is_tar, err := is_tarball(uri)

		if is_tar { return REMOTE_TAR, err }
		return REMOTE_ZIP, err
	} else {
		if fileinfo.IsDir() { return LOCAL_DIRECTORY, nil }

		is_tar, err := is_tarball(uri)

		if is_tar { return LOCAL_TAR, err }
		return LOCAL_ZIP, nil
	}
}

func is_tarball(uri *string) (bool, error) {
	switch path.Ext(*uri) {
	case ".gz":
		return true, nil
	case ".zip":
		return false, nil
	default:
		return false, PathError{fmt.Sprintf("Unsupported extension in given string %v", *uri)}
	}
}
