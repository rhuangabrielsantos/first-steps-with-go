package main

import "fmt"

const prefixHello = "Hello, "

func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}

	return prefixHello + name
}

func main() {
	fmt.Println(HelloWorld("world"))
}
