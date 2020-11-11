package parser

import (
	"io/ioutil"
	"testing"
)

func TestPraseCityList(t *testing.T) {
	contents, _ := ioutil.ReadFile("citylist_test_data.html")
	result := PraseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Items))
	}
}
