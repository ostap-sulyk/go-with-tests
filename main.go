package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}
func Addition(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Hello("world"))
}
