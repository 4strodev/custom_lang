package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	// Constant values
	NUMBER
	STRING
	IDENTIFIER

	// Grouping operators
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	// Assignment operators
	ASSIGNMENT
	PLUS_EQUALS
	MINUS_EQUALS

	// Logic operators
	GREATER
	GREATER_EQUALS
	LESS
	LESS_EQUALS
	OR
	AND
	EQUALS
	NOT
	NOT_EQUALS

	// Delimiters
	DOT
	COLON
	QUESTION
	COMMA
	SEMICOLON
	DOT_DOT

	// Arithmetic operators
	PLUS
	MINUS
	STAR
	PERCENT
	DASH
	PLUS_PLUS
	MINUS_MINUS

	// Reserved keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	FOR
	WHILE
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	Kind  TokenKind
	Value string
}

func TokenKindString(tokenKind TokenKind) string {
	switch tokenKind {
	case EOF:
		return "eof"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case EQUALS:
		return "equals"
	case NOT:
		return "not"
	case NOT_EQUALS:
		return "not_equals"
	case DOT:
		return "dot"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case SEMICOLON:
		return "semicolon"
	case DOT_DOT:
		return "dot_dot"
	case PLUS:
		return "plus"
	case MINUS:
		return "minus"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case DASH:
		return "slash"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case TYPEOF:
		return "typeof"
	case IN:
		return "in"
	}
	return ""
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, tokenKind := range expectedTokens {
		if token.Kind == tokenKind {
			return true
		}
	}

	return false
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind,
		value,
	}
}

func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}
