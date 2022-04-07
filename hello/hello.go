package main

import "fmt"

const (
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
