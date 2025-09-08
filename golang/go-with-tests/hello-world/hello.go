package helloworld

import (
	"fmt"

)

const (
	spanish = "spanish"
	french = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func main()  {
	fmt.Println(Hello("world", "english"))
	
}

func Hello(name string, language string)  string {
	if name == "" {
		name = "world"
	}
	return ReturnGreetingPrefix(language) + name
	
}

func ReturnGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	default:
		return englishHelloPrefix
		
	}
}
