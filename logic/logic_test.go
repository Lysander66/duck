package logic

import (
	"fmt"
	"testing"
)

func TestGenStruct(t *testing.T) {
	mysql := `root:123456@(127.0.0.1:3306)/world?charset=utf8&parseTime=True&loc=Local`
	model := GenStruct("world", "country", mysql)
	fmt.Println(model)
	// t.Log(model)
}
