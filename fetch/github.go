package fetch

import (
	"github.com/google/go-github/github"
	"time"
	"context"
	"net/url"
)

func get_github_repo_url(owner, repo string, branch *string) (*url.URL, *github.Response, error){
	d := time.Now().Add(20 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	client := github.NewClient(nil)

	var extra_opts github.RepositoryContentGetOptions
	if branch != nil {
		extra_opts = github.RepositoryContentGetOptions{
			Ref: *branch,
		}
	}

	return client.Repositories.GetArchiveLink(
		ctx,
		owner,
		repo,
		github.Zipball,
		&extra_opts,
	)
}
