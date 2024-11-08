package lgtmoon

import (
	"math/rand"
	"time"
)

var lgtmImages = []string{
    "https://lgtmoon.dev/lgtm01.png",
    "https://lgtmoon.dev/lgtm02.png",
    "https://lgtmoon.dev/lgtm03.png",
    // 必要に応じて他のLGTMOON画像のURLを追加
}

// GetRandomLgtmImageURL returns a random LGTM image URL from LGTMOON.
func GetRandomLgtmImageURL() string {
    rand.Seed(time.Now().Unix())
    return lgtmImages[rand.Intn(len(lgtmImages))]
}
