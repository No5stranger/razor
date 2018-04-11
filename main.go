package main

import (
	"os"
	"strings"

	"github.com/No5stranger/razor/alfred"
	"github.com/No5stranger/razor/shorteners"
)

func main() {
	if len(os.Args) >= 4 {
		shortener := os.Args[1]
		key := os.Args[2]
		longUrl := os.Args[3]
		if strings.Contains(longUrl, "$") {
			longUrl = strings.Replace(longUrl, "$", "", -1)
			if shortener == "goo.gl" {
				shortRes := shorteners.GenShortUrlByGoogl(key, longUrl)
				alfred.FormatXML(shortener, shortRes.ShortUrl)
			} else if shortener == "bitly" {
				shortRes := shorteners.GenShortUrlByBitly(key, longUrl)
				alfred.FormatXML(shortener, shortRes.ShortUrl)
			}
		} else {
			alfred.FormatXML(shortener, longUrl)
		}
	}
}
