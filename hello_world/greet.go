package main

import "fmt"

// Greet the user
func Greet(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Greet("Michael"))
}
