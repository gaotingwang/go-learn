package engine

import "github.com/gaotingwang/go-learn/crawler_distributed/config"

type Parser interface {
	Parser(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type ParserFunc func(content []byte, url string) ParseResult

type FuncParser struct {
	parser ParserFunc
	name   string
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{parser: parser, name: name}
}

func (fp *FuncParser) Parser(content []byte, url string) ParseResult {
	return fp.parser(content, url)
}

func (fp *FuncParser) Serialize() (name string, args interface{}) {
	return fp.name, nil
}

type Request struct {
	Url    string
	Parser Parser
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

type NilParser struct {
}

func (n NilParser) Parser(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (n NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}
