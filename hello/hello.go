package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const helloHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name

	case french:
		return helloHelloPrefix + name

	default:
		return englishHelloPrefix + name
	}

}

func main() {
	fmt.Println(Hello("world", "English"))
}
