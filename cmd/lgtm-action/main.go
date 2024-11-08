package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/db-r-hashimoto/lgtm_print/internal/lgtmoon"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func main() {
    token := os.Getenv("GITHUB_TOKEN")
    owner := os.Getenv("OWNER")
    repo := os.Getenv("REPO")
    issueNumberStr := os.Getenv("ISSUE_NUMBER")

    if token == "" {
        log.Fatal("GITHUB_TOKEN is not set")
    }
    if owner == "" || repo == "" {
        log.Fatal("OWNER or REPO is not set")
    }
    if issueNumberStr == "" {
        log.Fatal("ISSUE_NUMBER is not set")
    }

    // Issue番号を整数に変換
    issueNumber, err := strconv.Atoi(issueNumberStr)
    if err != nil {
        log.Fatalf("Invalid ISSUE_NUMBER: %v", err)
    }

    // OAuth2トークンを設定してGitHubクライアントを作成
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)

    // ランダムなLGTM画像URLを取得
    imageUrl, err := lgtmoon.GetRandomLgtmImageURL()
    if err != nil {
        log.Fatalf("Failed to fetch LGTM image: %v", err)
    }

    // コメント内容を作成
    comment := fmt.Sprintf("![LGTM](%s)", imageUrl)

    // GitHubのコメントAPIを使用してPRまたはIssueにコメントを投稿
    _, _, err = client.Issues.CreateComment(context.Background(), owner, repo, issueNumber, &github.IssueComment{
        Body: &comment,
    })
    if err != nil {
        log.Fatalf("Failed to post comment: %v", err)
    }

    log.Println("LGTM image posted successfully!")
}
