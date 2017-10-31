package main

import (
	"os"
	"strings"

	"github.com/No5stranger/razor/alfred"
	"github.com/No5stranger/razor/shorteners"
)

func main() {
	if len(os.Args) >= 4 {
		shortner := os.Args[1]
		key := os.Args[2]
		longUrl := os.Args[3]
		if strings.Contains(longUrl, "$") {
			longUrl = strings.Replace(longUrl, "$", "", -1)
			shortRes := shorteners.GenShortUrlByGoogl(key, longUrl)
			alfred.FormatXML(shortner, shortRes.ShortUrl)
		}
	}
}
