package parser

import (
	"GoPartice/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, _ := ioutil.ReadFile("profile_test_data.html")
	result := ParserProfile(contents, "三公子")
	if len(result.Items) != 1 {
		t.Errorf("Items 应该有1条结果，但是结果是：%d", len(result.Items))
	}
	profileInfo := result.Items[0]
	expected := model.Profile{
		Name:       "三公子",
		Gender:     "男",
		Age:        46,
		Height:     170,
		Weight:     55,
		Income:     "5-8千",
		Marriage:   "未婚",
		Education:  "高中及以下",
		Occupation: "其他职业",
		Hukou:      "浙江金华",
		Xingzuo:    "魔羯座(12.22-01.19)",
		House:      "住在单位宿舍",
		Car:        "未买车",
		Workplace:  "杭州上城区",
	}
	if profileInfo != expected {
		t.Errorf("期盼结果：%v,得到的结果：%v", expected, profileInfo)
	}
}
