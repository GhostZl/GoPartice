package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPraseCity(t *testing.T) {
	contents, _ := ioutil.ReadFile("city_test_data.html")
	result := ParserCity(contents)
	fmt.Println(result)
}
