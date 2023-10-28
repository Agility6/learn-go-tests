package main

import (
	"fmt"
)

const (
	french  = "French"
	spanish = "Spanish"

	helloPrefix       = "Hello, "
	spanisHelloPrefix = "Hola, "
	frenchHelloPrefix = "French "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// function private
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanisHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println("Hello, World")
}
