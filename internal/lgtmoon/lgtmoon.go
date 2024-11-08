package lgtmoon

import (
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

type LGTMImage struct {
    ID       int    `json:"id"`
    ImageURL string `json:"imageUrl"`
}


// GetRandomLgtmImageURL fetches a list of LGTM images from the API and returns a random image URL
func GetRandomLgtmImageURL() (string, error) {
    // APIエンドポイントのURL
    url := "https://lgtmeow.com/api/lgtm-images"

    // HTTPリクエストを送信
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // レスポンスが成功しているか確認
    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to fetch images from LGTMeow API")
    }

    // レスポンスボディをJSONデコード
    var images []LGTMImage
    if err := json.NewDecoder(resp.Body).Decode(&images); err != nil {
        return "", err
    }

    // 画像が存在するかチェック
    if len(images) == 0 {
        return "", errors.New("no images found")
    }

    // ランダムな画像を選択
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomIndex := rnd.Intn(len(images))
    return images[randomIndex].ImageURL, nil
}
