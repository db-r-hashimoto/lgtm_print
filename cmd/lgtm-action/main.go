package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/db-r-hashimoto/lgtm_print/internal/lgtmoon"
	"github.com/google/go-github/v48/github"
)

func main() {
    // GitHubトークンの取得
    token := os.Getenv("GITHUB_TOKEN")
    if token == "" {
        log.Fatal("GITHUB_TOKEN is not set")
    }

    // GitHubクライアントの初期化
    client := github.NewClient(nil)

    // ランダムなLGTM画像URLを取得
    imageUrl := lgtmoon.GetRandomLgtmImageURL()

    // コメント内容を作成
    comment := fmt.Sprintf("![LGTM](%s)", imageUrl)

    // GitHubのコメントAPIを使用してPRまたはIssueにコメントを投稿
    _, _, err := client.Issues.CreateComment(context.Background(), "owner", "repo", 1, &github.IssueComment{
        Body: &comment,
    })
    if err != nil {
        log.Fatalf("Failed to post comment: %v", err)
    }

    log.Println("LGTM image posted successfully!")
}
