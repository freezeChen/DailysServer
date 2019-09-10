package main

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	RpcAddReSs string
}

func main() {
	var jsonStr = `{
	"rPcaDdress":"address is "
}`

	var tag = new(Tag)
	json.Unmarshal([]byte(jsonStr), tag)

	fmt.Println(tag)
}
