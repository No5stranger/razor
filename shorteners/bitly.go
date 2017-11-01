package shorteners

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var BITLY_API_URL = "https://api-ssl.bitly.com/v3/shorten?access_token=%s&longUrl=%s"

type Bitly struct {
	GlobalHash string `json:"global_hash"`
	Hash       string `json:"hash"`
	LongUrl    string `json:"long_url"`
	NewHash    int    `json:"new_hash"`
	ShortUrl   string `json:"url"`
}

type RespBody struct {
	Data       Bitly `json:"data"`
	StatusCode int   `json:"status_code"`
}

func GenShortUrlByBitly(key string, shortUrl string) Bitly {
	escapedUrl := url.QueryEscape(shortUrl)
	apiUrl := fmt.Sprintf(BITLY_API_URL, key, escapedUrl)
	resp, err := http.Get(apiUrl)
	if err != nil {
		panic(err)
	}
	respBody := &RespBody{}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		panic(err)
	}
	return respBody.Data
}
