package fetcher

import (
	"fmt"
	"testing"
	"time"
)

func TestFetch(t *testing.T) {
	for i := 0; i < 100; i++ {
		res, err := Fetch("https://hanyu.baidu.com/s?wd=%E4%B8%BA&cf=rcmd&t=img&ptype=zici")
		time.Sleep(time.Millisecond * 100)
		fmt.Println(string(res), err)
	}
}
