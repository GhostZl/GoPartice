package parser

import (
	"GoPartice/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`

func ParserCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParserProfile(contents, name)
			},
		})
	}
	return result
}
