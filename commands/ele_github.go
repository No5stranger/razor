package commands

import (
	"fmt"
	"github.com/No5stranger/razor/alfred"
	"github.com/urfave/cli"
)

var FormateGithubEnterpriseCommand = cli.Command{
	Name:   "github_enterprise",
	Usage:  "format github enterprise project url",
	Action: FormatGithubEnterpirseUrl,
}

var host = "http://github.elenet.me"

var shortName = map[string]string{
	"a":  "alaya",
	"b":  "booking",
	"bo": "arsenal",
	"eu": "eus",
}

func FormatGithubEnterpirseUrl(c *cli.Context) error {
	var title = "open eleme github enterpise project"
	var targetURL string
	baseURL := "%s/waimai/%s"
	projectName := c.Args().First()
	if proj, ok := shortName[projectName]; ok {
		targetURL = fmt.Sprintf(baseURL, host, proj)
		alfred.FormatWebXML(title, targetURL)
	} else {
		targetURL = fmt.Sprintf(baseURL, host, projectName)
		alfred.FormatWebXML(title, targetURL)
	}
	return nil
}
