package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/db-r-hashimoto/lgtm_print/internal/lgtmoon"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

func main() {
    token := os.Getenv("GITHUB_TOKEN")
    owner := os.Getenv("OWNER")
    repo := os.Getenv("REPO")
    commentBody := os.Getenv("COMMENT_BODY")
    commentIDStr := os.Getenv("COMMENT_ID")

    if token == "" {
        log.Fatal("GITHUB_TOKEN is not set")
    }
    if owner == "" || repo == "" {
        log.Fatal("OWNER or REPO is not set")
    }

    // コメントIDを整数に変換
    commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
    if err != nil {
        log.Fatalf("Invalid COMMENT_ID: %v", err)
    }

    // OAuth2トークンを設定してGitHubクライアントを作成
    ctx := context.Background()
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: token},
    )
    tc := oauth2.NewClient(ctx, ts)
    client := github.NewClient(tc)

    // 画像が既に貼られているか確認
    if strings.Contains(commentBody, "![LGTM]") {
        log.Println("LGTM image already present in the comment.")
        return
    }

    // ランダムなLGTM画像URLを取得
    imageUrl, err := lgtmoon.GetRandomLgtmImageURL()
    if err != nil {
        log.Fatalf("Failed to fetch LGTM image: %v", err)
    }

    // GitHubのコメントAPIを使用してPRまたはIssueにコメントを投稿
    updatedComment := fmt.Sprintf("%s\n\n![LGTM](%s)", commentBody, imageUrl)

    _, _, err = client.Issues.EditComment(ctx, owner, repo, commentID, &github.IssueComment{
        Body: &updatedComment,
    })
    if err != nil {
        log.Fatalf("Failed to edit comment: %v", err)
    }


    log.Println("LGTM image posted successfully!")
}
