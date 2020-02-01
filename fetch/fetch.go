package fetch

type Option struct {
	File *File
	Err error
}

type File struct {
	Name string
	Contents []byte
}

func GithubRepo(user, repo string, branch *string, pipeline chan Option) {
	url, _, err := get_github_repo_url(user, repo, branch)

	if err != nil {
		pipeline <- Option{File: nil, Err: err}
		return
	}

	urlstring := url.String()
	RemoteRepoZip(&urlstring, pipeline)
}

func RemoteRepoTar(url *string, pipeline chan Option) {
	remote_download_and_read(url, pipeline, crawl_tar)
}

func RemoteRepoZip(url *string, pipeline chan Option) {
	remote_download_and_read(url, pipeline, crawl_zip)
}

func LocalRepoZip(path *string, pipeline chan Option) {
	defer close(pipeline)

	if err := crawl_zip(path, pipeline); err != nil {
		pipeline <- Option{
			File: nil,
			Err: err,
		}
	}
}

func LocalRepoTar(path *string, pipeline chan Option) {
	defer close(pipeline)

	if err := crawl_tar(path, pipeline); err != nil {
		pipeline <- Option{
			File: nil,
			Err: err,
		}
	}
}

func LocalRepoDir(path *string, pipeline chan Option) {
	defer close(pipeline)

	if err := crawl_dir(path, pipeline); err != nil {
		pipeline <- Option{
			File: nil,
			Err: err,
		}
	}
}

func remote_download_and_read(url *string, pipeline chan Option, lambda Crawlable) {
	defer close(pipeline)

	tmpfile, err := download_to_tmp(url)
	if err != nil {
		pipeline <- Option{File: nil, Err: err}
	}

	if err := lambda(&tmpfile, pipeline); err != nil {
		pipeline <- Option{File: nil, Err: err,}
	}
}
