package alfred

import (
	"encoding/xml"
	"os"
)

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

func FormatXML(shortener string, url string) {
	items := &AlfredItems{Version: "1"}
	title := "short url by " + shortener
	items.Item = append(items.Item, AlfredItem{Arg: url, Title: title, SubTitle: url})
	output, err := xml.Marshal(items)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
