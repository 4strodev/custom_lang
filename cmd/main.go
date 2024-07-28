package main

import (
	"os"

	"github.com/4strodev/custom_lang/pkg/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.lang")
	source := string(bytes)

	tokens := lexer.Tokenize(source)

	for _, token := range tokens {
		token.Debug()
	}
}
