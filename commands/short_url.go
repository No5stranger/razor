package commands

import (
	"strings"

	"github.com/No5stranger/razor/alfred"
	"github.com/No5stranger/razor/shorteners"

	"github.com/urfave/cli"
)

var ShortUrlCommand = cli.Command{
	Name:   "shorter",
	Usage:  "shorter key url$",
	Action: ShortUrl,
}

var apiKey string

var ShorterKeyFlag = cli.StringFlag{
	Name:        "key",
	Usage:       "shorter api key",
	Destination: &apiKey,
}

func ShortUrl(c *cli.Context) error {
	url := c.Args().First()
	if strings.Contains(url, "$") {
		url = strings.Replace(url, "$", "", -1)
		shorterRep := shorteners.GenShortUrlByBitly(apiKey, url)
		alfred.FormatXML("bitly", shorterRep.ShortUrl)
	} else {
		alfred.FormatXML("bitly", url)
	}
	return nil
}
