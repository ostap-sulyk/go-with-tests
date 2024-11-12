package main

const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	franchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixGreeting(language) + name
}

func prefixGreeting(language string) (prefix string) {
	switch language {
	case french:
		prefix = french
	case spanish:
		prefix = spanish
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	println(Hello("world", ""))
}
