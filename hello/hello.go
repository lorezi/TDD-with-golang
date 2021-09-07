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

	switch language {

	case spanish:
		return spanishHelloPrefix + name

	case french:
		return frenchHelloPrefix + name

	case pidgin:
		return pidginHelloPrefix + name

	default:
		return englishHelloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("world", "English"))
}
