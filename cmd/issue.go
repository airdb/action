/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v32/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "github action issue and issue_comment operation",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		githubEventPath := os.Getenv("GITHUB_EVENT_PATH")
		githubEventName := os.Getenv("GITHUB_EVENT_NAME")
		githubToken := os.Getenv("GITHUB_TOKEN")

		if githubEventPath == "" {
			fmt.Println("GITHUB_EVENT_PATH is null.")
			return
		}

		var ghIssueComment GithubIssueComment

		content, err := ioutil.ReadFile(githubEventPath)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(content, &ghIssueComment)
		if err != nil {
			panic(err)
		}

		SetIssueLabel(githubToken,
			ghIssueComment.Repository.Owner.Login,
			ghIssueComment.Repository.Name,
			ghIssueComment.Issue.ID,
			ghIssueComment.Comment.Body,
		)

		fmt.Printf("event: %s, file: %s\n", githubEventName, githubEventPath)
		fmt.Printf("action: %s, issue title: %s, issue body: %s, issue user: %s\n",
			ghIssueComment.Action,
			ghIssueComment.Issue.Title,
			ghIssueComment.Issue.Body,
			ghIssueComment.Issue.User.Login,
		)

		fmt.Printf("comment user: %s, comment body: %s\n",
			ghIssueComment.Comment.User.Login,
			ghIssueComment.Comment.Body,
		)
	},
}

func issueCmdInit() {
	rootCmd.AddCommand(issueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:// issueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type GithubIssueComment struct {
	Action  string `json:"action"`
	Comment struct {
		AuthorAssociation string `json:"author_association"`
		Body              string `json:"body"`
		CreatedAt         string `json:"created_at"`
		HTMLURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		IssueURL          string `json:"issue_url"`
		NodeID            string `json:"node_id"`
		UpdatedAt         string `json:"updated_at"`
		URL               string `json:"url"`
		User              struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"user"`
	} `json:"comment"`
	Issue struct {
		ActiveLockReason  interface{}   `json:"active_lock_reason"`
		Assignee          interface{}   `json:"assignee"`
		Assignees         []interface{} `json:"assignees"`
		AuthorAssociation string        `json:"author_association"`
		Body              string        `json:"body"`
		ClosedAt          interface{}   `json:"closed_at"`
		Comments          int64         `json:"comments"`
		CommentsURL       string        `json:"comments_url"`
		CreatedAt         string        `json:"created_at"`
		EventsURL         string        `json:"events_url"`
		HTMLURL           string        `json:"html_url"`
		ID                int           `json:"id"`
		Labels            []interface{} `json:"labels"`
		LabelsURL         string        `json:"labels_url"`
		Locked            bool          `json:"locked"`
		Milestone         interface{}   `json:"milestone"`
		NodeID            string        `json:"node_id"`
		Number            int64         `json:"number"`
		RepositoryURL     string        `json:"repository_url"`
		State             string        `json:"state"`
		Title             string        `json:"title"`
		UpdatedAt         string        `json:"updated_at"`
		URL               string        `json:"url"`
		User              struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"user"`
	} `json:"issue"`
	Organization struct {
		AvatarURL        string `json:"avatar_url"`
		Description      string `json:"description"`
		EventsURL        string `json:"events_url"`
		HooksURL         string `json:"hooks_url"`
		ID               int64  `json:"id"`
		IssuesURL        string `json:"issues_url"`
		Login            string `json:"login"`
		MembersURL       string `json:"members_url"`
		NodeID           string `json:"node_id"`
		PublicMembersURL string `json:"public_members_url"`
		ReposURL         string `json:"repos_url"`
		URL              string `json:"url"`
	} `json:"organization"`
	Repository struct {
		ArchiveURL       string      `json:"archive_url"`
		Archived         bool        `json:"archived"`
		AssigneesURL     string      `json:"assignees_url"`
		BlobsURL         string      `json:"blobs_url"`
		BranchesURL      string      `json:"branches_url"`
		CloneURL         string      `json:"clone_url"`
		CollaboratorsURL string      `json:"collaborators_url"`
		CommentsURL      string      `json:"comments_url"`
		CommitsURL       string      `json:"commits_url"`
		CompareURL       string      `json:"compare_url"`
		ContentsURL      string      `json:"contents_url"`
		ContributorsURL  string      `json:"contributors_url"`
		CreatedAt        string      `json:"created_at"`
		DefaultBranch    string      `json:"default_branch"`
		DeploymentsURL   string      `json:"deployments_url"`
		Description      interface{} `json:"description"`
		Disabled         bool        `json:"disabled"`
		DownloadsURL     string      `json:"downloads_url"`
		EventsURL        string      `json:"events_url"`
		Fork             bool        `json:"fork"`
		Forks            int64       `json:"forks"`
		ForksCount       int64       `json:"forks_count"`
		ForksURL         string      `json:"forks_url"`
		FullName         string      `json:"full_name"`
		GitCommitsURL    string      `json:"git_commits_url"`
		GitRefsURL       string      `json:"git_refs_url"`
		GitTagsURL       string      `json:"git_tags_url"`
		GitURL           string      `json:"git_url"`
		HasDownloads     bool        `json:"has_downloads"`
		HasIssues        bool        `json:"has_issues"`
		HasPages         bool        `json:"has_pages"`
		HasProjects      bool        `json:"has_projects"`
		HasWiki          bool        `json:"has_wiki"`
		Homepage         interface{} `json:"homepage"`
		HooksURL         string      `json:"hooks_url"`
		HTMLURL          string      `json:"html_url"`
		ID               int64       `json:"id"`
		IssueCommentURL  string      `json:"issue_comment_url"`
		IssueEventsURL   string      `json:"issue_events_url"`
		IssuesURL        string      `json:"issues_url"`
		KeysURL          string      `json:"keys_url"`
		LabelsURL        string      `json:"labels_url"`
		Language         interface{} `json:"language"`
		LanguagesURL     string      `json:"languages_url"`
		License          struct {
			Key    string `json:"key"`
			Name   string `json:"name"`
			NodeID string `json:"node_id"`
			SpdxID string `json:"spdx_id"`
			URL    string `json:"url"`
		} `json:"license"`
		MergesURL        string      `json:"merges_url"`
		MilestonesURL    string      `json:"milestones_url"`
		MirrorURL        interface{} `json:"mirror_url"`
		Name             string      `json:"name"`
		NodeID           string      `json:"node_id"`
		NotificationsURL string      `json:"notifications_url"`
		OpenIssues       int64       `json:"open_issues"`
		OpenIssuesCount  int64       `json:"open_issues_count"`
		Owner            struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"owner"`
		Private         bool   `json:"private"`
		PullsURL        string `json:"pulls_url"`
		PushedAt        string `json:"pushed_at"`
		ReleasesURL     string `json:"releases_url"`
		Size            int64  `json:"size"`
		SSHURL          string `json:"ssh_url"`
		StargazersCount int64  `json:"stargazers_count"`
		StargazersURL   string `json:"stargazers_url"`
		StatusesURL     string `json:"statuses_url"`
		SubscribersURL  string `json:"subscribers_url"`
		SubscriptionURL string `json:"subscription_url"`
		SvnURL          string `json:"svn_url"`
		TagsURL         string `json:"tags_url"`
		TeamsURL        string `json:"teams_url"`
		TreesURL        string `json:"trees_url"`
		UpdatedAt       string `json:"updated_at"`
		URL             string `json:"url"`
		Watchers        int64  `json:"watchers"`
		WatchersCount   int64  `json:"watchers_count"`
	} `json:"repository"`
	Sender struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		NodeID            string `json:"node_id"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"sender"`
}

func SetIssueLabel(accessToken, owner, repo string, issueNumber int, label string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	labels, resp, err := client.Issues.AddLabelsToIssue(ctx, owner, repo, issueNumber, []string{label})
	fmt.Println("resp", labels, resp, err)
}
