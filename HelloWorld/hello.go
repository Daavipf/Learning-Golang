package main

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "
const helloPortuguesePrefix = "Olá, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	case portuguese:
		prefix = helloPortuguesePrefix
	default:
		prefix = helloEnglishPrefix
	}

	return
}
