package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch {
	case language == spanish:
		prefix = spanishHelloPrefix
	case language == french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("World", ""))
}
