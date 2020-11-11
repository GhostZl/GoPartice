package parser

import (
	"GoPartice/crawler/engine"
	"regexp"
)

const host = "http://newcar.xcar.com.cn"

var carModelRe = regexp.MustCompile(`<a href="(/\d+/)" target="_blank" class="list_img">`)
var carListRe = regexp.MustCompile(`<a href="(//newcar.xcar.com.cn/car/[\d+-]+\d+/)"`)

func ParseCarList(
	contents []byte, url string) engine.ParseResult {
	matches := carModelRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       host + string(m[1]),
				ParseFunc: ParseCarModel,
			})
	}

	matches = carListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       "http:" + string(m[1]),
				ParseFunc: ParseCarList,
			})
	}

	return result
}
