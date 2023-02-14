package main

import (
	"encoding/json"
	"fmt"
	"github.com/gaotingwang/go-learn/jsonrpc/client"
	"github.com/gaotingwang/go-learn/jsonrpc/server"
	"log"
	"time"
)

func main() {
	// 服务端
	const host = ":1234"
	go server.Rpc(host, server.Service{})
	time.Sleep(time.Second) //保证服务起来

	// 客户端
	user, err := client.CreateProcessor(host)
	if err != nil {
		panic(err)
	}

	// rpc 接口调用返回结果，对json反序列化时，若对象属性中有interface{}类型，会映射为map[string]interface{}
	fmt.Printf("%+v\n", user)
	if m, ok := user.Favorite.(map[string]interface{}); ok {
		for k, v := range m {
			switch k {
			case "Desc":
				fmt.Println(k, ":", v)
				desc := *decodeByteArray(v)
				fmt.Println(k, ":", desc)
			default:
				fmt.Println(k, ":", v)
			}
		}
		//fmt.Printf("%s\n", m["Desc"].([]byte))
		//if s, ok := m["Type"].(string); ok {
		//	fmt.Printf("%s\n", s)
		//}
	}

	//bolB, _ := json.Marshal([]byte("This is man"))
	//var a []byte
	//json.Unmarshal(bolB, &a)
	//fmt.Println(a)
}

func decodeByteArray(v interface{}) *[]byte {
	var desc []byte
	err := json.Unmarshal([]byte("\""+v.(string)+"\""), &desc)
	if err != nil {
		log.Printf("error decoding : %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
	}
	return &desc
}
