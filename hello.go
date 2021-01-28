package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishGoodbyePrefix = "Goodbye, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Goodbye(name string) string {
	if name == "" {
		name = "Cruel World"
	}
	return englishGoodbyePrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
	fmt.Println(Hello("Chris"))
}
