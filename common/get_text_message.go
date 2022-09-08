package common

import "github.com/xanzy/go-gitlab"

func GetTextMessage(commit gitlab.Commit) string {
	return "[New commit to master]: \nCommit message: \"" + commit.Title + "\"\nCommit author: " + commit.AuthorName + "\nCommit time: " + commit.CommittedDate.Format("02-01-2006 15:04:05") + "\nCommit url: " + commit.WebURL + "\n\n🚀 Bot have build and deploy new version of project. You can view it at: https://beta.cronmail.net"
}
