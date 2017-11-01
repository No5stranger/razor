package shorteners

import "encoding/json"
import "net/http"
import "bytes"

var GOOGL_API_URL = "https://www.googleapis.com/urlshortener/v1/url?key="

type Googl struct {
	ShortUrl string `json:"id"`
	Kind     string `json:"kind"`
	LongUrl  string `json:"longUrl"`
}

type ReqBody struct {
	LongUrl string `json:"longUrl"`
}

func GenShortUrlByGoogl(key string, url string) *Googl {
	var reqBody ReqBody
	reqBody.LongUrl = url
	jsonBody, _ := json.Marshal(reqBody)
	apiUrl := GOOGL_API_URL + key
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	googl := &Googl{}
	err = json.NewDecoder(resp.Body).Decode(googl)
	if err != nil {
		panic(err)
	}
	return googl
}
