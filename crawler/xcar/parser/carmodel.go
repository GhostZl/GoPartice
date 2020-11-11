package parser

import (
	"GoPartice/crawler/engine"
	"regexp"
)

var carDetailRe = regexp.MustCompile(`<a href="(/m\d+/)" target="_blank"`)

func ParseCarModel(
	contents []byte, url string) engine.ParseResult {
	matches := carDetailRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		url := host + string(m[1])
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       url,
				ParseFunc: ParseCarDetail,
			})
	}

	return result
}
