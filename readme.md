## Razor
a command line tool to get short url by goo.gl or bitly


### Support Shortener

#### goo.gl
* [get start](https://developers.google.com/url-shortener/v1/getting_started)
* create api key for [Url Shortener] in [Google console](https://console.developers.google.com/apis/credentials)

#### bitly
* create a bitly account in [bitly](https://bityly.com)
* generate access_token in settings


### Usage
```shell
go install github.com/No5stranger/razor
$GOPATH/bin/razor {shortener} {auth_key} {url}+$
```
