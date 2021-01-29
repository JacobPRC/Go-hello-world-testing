package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishGoodbyePrefix = "Goodbye, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanishGoodbyePrefix = "Adios, "
const frenchGoodbyePrefix = "Au revoir, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func Goodbye(name string, language string) string {
	if name == "" {
		name = "Cruel World"
	}

	prefix := englishGoodbyePrefix

	switch language {
	case "French":
		prefix = frenchGoodbyePrefix
	case "Spanish":
		prefix = spanishGoodbyePrefix
	}

	return prefix + name

}

func main() {
	fmt.Println(Hello("Chris", ""))
	fmt.Println(Hello("Chris", ""))
}
