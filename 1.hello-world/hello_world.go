package main

import "fmt"

const enflishHelloPrefix = "Hello, "

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const englishHelloPrefix = "Bonjour, "

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = enflishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(HelloWorld("world", ""))
}
