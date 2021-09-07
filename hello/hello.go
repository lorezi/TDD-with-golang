package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const pidgin = "Pidgin"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const pidginHelloPrefix = "Afa, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {

	case spanish:
		prefix = spanishHelloPrefix
		return prefix

	case french:
		prefix = frenchHelloPrefix
		return prefix

	case pidgin:
		prefix = pidginHelloPrefix
		return prefix

	default:
		prefix = englishHelloPrefix
		return prefix
	}

}

func main() {
	fmt.Println(Hello("world", "English"))
}
