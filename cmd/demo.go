package main

import (
	"example/apisecrets"
	"fmt"
)

func main() {
	v := apisecrets.Memory("my-fake-key")

	err := v.Set("demo-key", "demo-value")
	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}

	fmt.Println("Plain: ", plain)
	
}