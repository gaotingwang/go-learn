package worker

import (
	"errors"
	"fmt"
	"log"

	"github.com/gaotingwang/go-learn/crawler/engine"
	"github.com/gaotingwang/go-learn/crawler/zhenai/parser"
	"github.com/gaotingwang/go-learn/crawler_distributed/config"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

func SerializedRequest(request engine.Request) Request {
	name, args := request.Parser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializedResult(pr engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: pr.Items,
	}

	for _, r := range pr.Requests {
		result.Requests = append(result.Requests, SerializedRequest(r))
	}

	return result
}

func DeserializedRequest(request Request) (engine.Request, error) {
	par, err := deserializedParser(request.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    request.Url,
		Parser: par,
	}, nil
}

func DeserializedResult(pr ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: pr.Items,
	}

	for _, r := range pr.Requests {
		engineReq, err := DeserializedRequest(r)
		if err != nil {
			log.Printf("error deserilizing request:%v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializedParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		// log.Print("ParseCityList works well")
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		// log.Print("ParseCity works well")
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if profileParser, ok := p.Args.(map[string]interface{}); ok {
			// log.Print(userName)
			return parser.NewProfileParser(profileParser["UserUrl"].(string), profileParser["Name"].(string), profileParser["Info"].([]byte)), nil
		} else {
			// log.Print("ParseProfile falls")
			return nil, fmt.Errorf("invalid arg:%v", p.Args)
		}
	default:
		// log.Print(p.Name) // 很关键的Debug
		return nil, errors.New("unknown parser name")
	}
}
