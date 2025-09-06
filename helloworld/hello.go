package helloworld

const (
	spanish = "Spanish"
	french  = "French"
	nepali  = "Nepali"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	nepaliHelloPrefix  = "Namaste, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "Club"
	}
	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case nepali:
		prefix = nepaliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
