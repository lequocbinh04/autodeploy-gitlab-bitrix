package common

import (
	"github.com/hashicorp/go-retryablehttp"
	"github.com/xanzy/go-gitlab"
	"os"
)

func GetCommit() (gitlab.Commit, error) {
	git, err := gitlab.NewClient(os.Getenv("GITLAB_TOKEN"))
	if err != nil {
		return gitlab.Commit{}, err
	}
	repo, _, err := git.Commits.GetCommit(os.Getenv("PID"), "master", func(request *retryablehttp.Request) error {
		return nil
	})
	return *repo, nil
}
