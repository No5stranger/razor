package commands

import (
	"fmt"
	"github.com/No5stranger/razor/alfred"
	"github.com/urfave/cli"
)

var traceHost = "http://trace.elenet.me"

var FormatTraceUrlCommand = cli.Command{
	Name:   "trace",
	Usage:  "open project trace page",
	Action: FormatTraceURL,
}

var FormatSearchTraceCommand = cli.Command{
	Name:   "trace_search",
	Usage:  "search trace by request id",
	Action: FormatTraceSearchURL,
}

var shortProjName = map[string]string{
	"b":    "biz.booking",
	"bo":   "biz.bos",
	"eu":   "zeus.eus",
	"goeu": "biz.goeus",
	"ugc":  "biz.ugc",
}

func FormatTraceURL(c *cli.Context) error {
	var title = "open project trace"
	baseURL := "%s/#/application?appId=%s"
	var targetURL string
	projectName := c.Args().First()
	if proj, ok := shortProjName[projectName]; ok {
		targetURL := fmt.Sprintf(baseURL, traceHost, proj)
		alfred.FormatWebXML(title, targetURL)
	} else {
		targetURL = fmt.Sprintf(baseURL, traceHost, projectName)
		alfred.FormatWebXML(title, targetURL)
	}
	return nil
}

func FormatTraceSearchURL(c *cli.Context) error {
	var title = "search trace by request id"
	baseURL := "%s/#/search?type=3&id=%s"
	var targetURL string
	reqID := c.Args().First()
	targetURL = fmt.Sprintf(baseURL, traceHost, reqID)
	alfred.FormatWebXML(title, targetURL)
	return nil
}
