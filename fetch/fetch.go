package fetch

type File struct {
	Name string
	Contents []byte
}

func RemoteRepoTar(url *string, files chan File, errors chan error) {
	if tmpfile := remote_download(url, files, errors); tmpfile != nil {
		LocalRepoTar(tmpfile, files, errors)
	}
}

func RemoteRepoZip(url *string, files chan File, errors chan error) {
	if tmpfile := remote_download(url, files, errors); tmpfile != nil {
		LocalRepoZip(tmpfile, files, errors)
	}
}

func LocalRepoZip(path *string, files chan File, errors chan error) {
	crawl_zip(path, files, errors)
	close_pipes(files, errors)
}

func LocalRepoTar(path *string, files chan File, errors chan error) {
	crawl_tar(path, files, errors)
	close_pipes(files, errors)
}

func LocalRepoDir(path *string, files chan File, errors chan error) {
	crawl_dir(path, files, errors)
	close_pipes(files, errors)
}

func remote_download(url *string, files chan File, errors chan error) *string {
	if tmpfile, err := download_to_tmp(url); err != nil {
		errors <- err
		close_pipes(files, errors)
		return nil
	} else {
		return &tmpfile
	}
}

func close_pipes(files chan File, errors chan error) {
		close(files)
		close(errors)
}
