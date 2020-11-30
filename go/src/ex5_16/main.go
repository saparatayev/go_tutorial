package main

import (
	"./strings"
	"fmt"
)

func main() {
	strs := []string{"aaa","bbb","CCC","555"}
	fmt.Println(strings.Join("++",strs...))
}