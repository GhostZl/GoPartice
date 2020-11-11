package parser

import (
	"GoPartice/crawler/engine"
	"GoPartice/crawler/model"
	"regexp"
	"strconv"
	"strings"
)

const basicInfoRe = `<div class="m-btn purple"[^>]+>([^<]+)</div>`
const otherInfoRe = `<div class="m-btn pink"[^>]+>([^<]+)</div>`
const genderRe = `"genderString":"([^\"])."`

func ParserProfile(contents []byte, name string) engine.ParseResult {
	re := regexp.MustCompile(basicInfoRe)
	basicMathces := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	item := model.Profile{}
	//基础信息
	for index, m := range basicMathces {
		switch index {
		case 0:
			item.Marriage = string(m[1])
		case 1:
			item.Age, _ = strconv.Atoi(string(m[1][:len(m[1])-3]))
		case 2:
			item.Xingzuo = string(m[1])
		case 3:
			item.Height, _ = strconv.Atoi(string(m[1][:len(m[1])-2]))
		case 4:
			item.Weight, _ = strconv.Atoi(string(m[1][:len(m[1])-2]))
		case 5:
			if strings.Index(string(m[1]), ":") >= 0 {
				rec := regexp.MustCompile(`[^:]+:(.*)`)
				va := rec.FindSubmatch(m[1])
				item.Workplace = string(va[1])
			} else {
				item.Workplace = string(m[1])
			}
		case 6:
			if strings.Index(string(m[1]), ":") >= 0 {
				rec := regexp.MustCompile(`[^:]+:(.*)`)
				va := rec.FindSubmatch(m[1])
				item.Income = string(va[1])
			} else {
				item.Income = string(m[1])
			}
		case 7:
			item.Occupation = string(m[1])
		case 8:
			item.Education = string(m[1])
		}
	}
	re = regexp.MustCompile(otherInfoRe)
	otherMathces := re.FindAllSubmatch(contents, -1)
	if otherMathces != nil && len(otherMathces) >= 2 {
		infos := string(otherMathces[1][1])
		if strings.Index(infos, ":") >= 0 {
			rec := regexp.MustCompile(`[^:]+:(.*)`)
			va := rec.FindSubmatch(otherMathces[1][1])
			item.Hukou = string(va[1])
		} else {
			item.Hukou = infos
		}
	}
	if otherMathces != nil && len(otherMathces) >= 6 {
		infos := string(otherMathces[5][1])
		if strings.Index(infos, ":") >= 0 {
			rec := regexp.MustCompile(`[^:]+:(.*)`)
			va := rec.FindSubmatch(otherMathces[5][1])
			item.House = string(va[1])
		} else {
			item.House = infos
		}
	}
	if otherMathces != nil && len(otherMathces) >= 7 {
		infos := string(otherMathces[6][1])
		if strings.Index(infos, ":") >= 0 {
			rec := regexp.MustCompile(`[^:]+:(.*)`)
			va := rec.FindSubmatch(otherMathces[6][1])
			item.Car = string(va[1])
		} else {
			item.Car = infos
		}
	}
	re = regexp.MustCompile(genderRe)
	genderMatches := re.FindSubmatch(contents)
	if genderMatches != nil && len(genderMatches) > 0 {
		item.Gender = string(genderMatches[1])
	}
	item.Name = name
	result.Requests = make([]engine.Request, 0)
	result.Items = append(result.Items, item)
	return result
}
