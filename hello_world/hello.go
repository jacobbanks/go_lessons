package hello


import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
var language_map = map[string]string{
	"spanish": spanishHelloPrefix,
	"english": englishHelloPrefix,
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return language_map[language] + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
