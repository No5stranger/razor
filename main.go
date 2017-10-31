package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"os"
	"strings"
)

var GOOGL_API_URL = "https://www.googleapis.com/urlshortener/v1/url?key="

type ReqBody struct {
	LongUrl string `json:"longUrl"`
}

type RespBody struct {
	Kind     string `json:"kind"`
	ShortUrl string `json:"id"`
	LongUrl  string `json:"longUrl"`
}

type AlfredItems struct {
	Version string       `xml:"version,attr"`
	XMLName xml.Name     `xml:"items"`
	Item    []AlfredItem `xml:"item"`
}

type AlfredItem struct {
	Arg      string `xml:"arg,attr"`
	Title    string `xml:"title"`
	SubTitle string `xml:"subtitle"`
}

func formatXML(url string) {
	items := &AlfredItems{Version: "1"}
	items.Item = append(items.Item, AlfredItem{Arg: url, Title: "short url by goo.gl", SubTitle: url})
	output, err := xml.Marshal(items)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}

func genShortUrl(key string, url string) string {
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
	respBody := &RespBody{}
	err = json.NewDecoder(resp.Body).Decode(respBody)
	if err != nil {
		panic(err)
	}
	return respBody.ShortUrl
}

func main() {
	if len(os.Args) >= 3 {
		key := os.Args[1]
		longUrl := os.Args[2]
		if strings.Contains(longUrl, "$") {
			longUrl = strings.Replace(longUrl, "$", "", -1)
			shortUrl := genShortUrl(key, longUrl)
			formatXML(shortUrl)
		}
	}
}
