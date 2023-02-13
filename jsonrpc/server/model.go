package server

// 作为 rpc 接口的返回结果，属性首字母必须大写，来保证json可进行序列化与反序列化
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Content  []byte
	Favorite interface{}
}

func NewUser(name string, age int, content []byte, favorite interface{}) User {
	return User{Name: name, Age: age, Content: content, Favorite: favorite}
}

type Some struct {
	Name string
	Type string
}

func NewSome(name string, Type string) *Some {
	return &Some{Name: name, Type: Type}
}
