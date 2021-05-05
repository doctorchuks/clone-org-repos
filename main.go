package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)

	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)

	org := os.Getenv("GITHUB_ORG")

	if org == "" {
		log.Fatal("GITHUB_ORG environment variable is not set.")
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		CheckIfError(err)
	}

	variables := map[string]interface{}{
		"repoCursor":   (*githubv4.String)(nil),
		"organisation": githubv4.String(org),
	}

	var allNodes []node

	for {
		err := client.Query(context.Background(), &query, variables)

		if err != nil {
			CheckIfError(err)
		}

		allNodes = append(allNodes, query.Organization.Repositories.Nodes...)

		if !query.Organization.Repositories.PageInfo.HasNextPage {
			break
		}

		variables["repoCursor"] = githubv4.NewString(query.Organization.Repositories.PageInfo.EndCursor)

	}

	for _, node := range allNodes {
		Info("Cloning repository: " + node.Name)
		Info("SSH CLone URL: " + node.SshUrl)
		CloneOrUpdate(filepath.Join(homeDir, org, node.Name), node.SshUrl)
	}
}

var query struct {
	Organization struct {
		Repositories struct {
			Nodes    []node
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"repositories(first: 100, after: $repoCursor)"`
	} `graphql:"organization(login: $organisation)"`
}

type node struct {
	Name   string
	SshUrl string
}

func CloneOrUpdate(directory string, url string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		_, err := git.PlainClone(directory, false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		})
		CheckIfError(err)
	} else {
		// If the directory does exist open it and pull
		r, err := git.PlainOpen(directory)
		CheckIfError(err)

		w, err := r.Worktree()
		CheckIfError(err)

		Info("git pull origin")
		w.Pull(&git.PullOptions{RemoteName: "origin"})

		// Print the latest commit that was just pulled
		ref, err := r.Head()
		CheckIfError(err)

		commit, err := r.CommitObject(ref.Hash())
		CheckIfError(err)

		fmt.Println(commit)
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
