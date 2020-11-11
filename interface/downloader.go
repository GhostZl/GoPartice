package main

import (
	"GoPartice/interface/infra"
	"fmt"
)

func getRetriever() Retriever {
	return &infra.Retriever{}
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{
		"name":   "Ghost",
		"course": "Ghost",
	})
	fmt.Println()
}

func session(se RetrieverPoster) {
	se.Get("https://www.baidu.com")
}

type RetrieverPoster interface {
	Retriever
	Poster
}

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func main() {
	var r Retriever = getRetriever()
	fmt.Println(r.Get("https://www.baidu.com"))
}
