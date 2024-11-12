package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
func Addition(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Hello("world"))
}
