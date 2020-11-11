package engine

type Request struct {
	Url       string
	ParseFunc ParserFunc
}
type ParseResult struct {
	Requests []Request
	Items    []Item
}
type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParserFunc func(
	contents []byte, url string) ParseResult

func NilPraser([]byte) ParseResult {
	return ParseResult{}
}
