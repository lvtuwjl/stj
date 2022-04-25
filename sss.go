package sss

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cast"
)

// func main() {
// 	fmt.Println(123)
// }

type D struct {
	Name string
	Age  int
	Addr string
}

func StructToJson(d *D) string {
	bytes, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	var s string

	i := cast.ToInt(s)

	fmt.Println("i value111 is:", i)

	return string(bytes)
}
