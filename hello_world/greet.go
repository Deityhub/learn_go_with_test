package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

// Greet the user
func Greet(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	if lang == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Greet("Michael", "English"))
}
