package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Greet the user
func Greet(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Greet("Michael"))
}
