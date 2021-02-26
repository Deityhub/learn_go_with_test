package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
)

func greeting(prefix, name string) string {
	return prefix + name
}

func generatePrefix(lang string) (prefix string) { //default value for prefix is empty string ""
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
		break
	case french:
		prefix = frenchHelloPrefix
		break
	default:
		prefix = englishHelloPrefix
	}

	// another way of returning the prefix is by
	// just specifying return, prefix will be returned automatically
	return prefix
}

// Greet the user
func Greet(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := generatePrefix(lang)
	return greeting(prefix, name)
}

func main() {
	fmt.Println(Greet("Michael", "English"))
}
