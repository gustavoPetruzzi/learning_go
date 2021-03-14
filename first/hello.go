package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const spanish = "Spanish"
const french = "French"

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("crack", "english"))
}