package main

import(
	"unicode"
)

const (
	//Characters
	LETTER = iota
	DIGIT = iota
	WHITESPACE = iota

	//Tokens
	INT_LIT = iota
	IDENT = iota
	ASSIGN_OP = iota
	ADD_OP = iota
	SUB_OP = iota
	MULT_OP = iota
	DIV_OP = iota
	LEFT_PAREN = iota
	RIGHT_PAREN = iota
	EXPR_END = iota
	LEFT_BRAC = iota
	RIGHT_BRAC = iota
	GREAT_OP = iota
	LESS_OP = iota

	//Keywords
	IF_IDENT = iota
	FOR_IDENT = iota
	FUNC_IDENT = iota
	RETURN_IDENT = iota

	UNKNOWN = iota
)

const (
	STRIF = "if"
	STRFOR = "for"
	STRFUNC = "func"
	STRRETURN = "return"
)

func getKeyword(lexeme string) int {
	switch lexeme {
	case STRIF: return IF_IDENT
	case STRFOR: return FOR_IDENT
	case STRFUNC: return FUNC_IDENT
	case STRRETURN: return RETURN_IDENT
	}
	return IDENT
}

func getTokenString(ident int) string {
	switch ident {
	case LETTER: return "LETTER"
	case DIGIT: return "DIGIT"
	case WHITESPACE: return "WHITESPACE"
	case UNKNOWN: return "UNKNOWN"
	case INT_LIT: return "INT_LIT"
	case IDENT: return "IDENT"
	case ASSIGN_OP: return "ASSIGN_OP"
	case ADD_OP: return "ADD_OP"
	case SUB_OP: return "SUB_OP"
	case MULT_OP: return "MULT_OP"
	case DIV_OP: return "DIV_OP"
	case LEFT_PAREN: return "LEFT_PAREN"
	case RIGHT_PAREN: return "RIGHT_PAREN"
	case EXPR_END: return "EXPR_END"
	case LEFT_BRAC: return "LEFT_BRAC"
	case RIGHT_BRAC: return "RIGHT_BRAC"
	case GREAT_OP: return "GREAT_OP"
	case LESS_OP: return "LESS_OP"
	case IF_IDENT: return "IF_IDENT"
	case FOR_IDENT: return "FOR_IDENT"
	case FUNC_IDENT: return "FUNC_IDENT"
	case RETURN_IDENT: return "RETURN_IDENT"
	}
	return "INVALID_IDENT"
}

func charClass(char rune) int {
	switch {
	case unicode.IsDigit(char): return DIGIT
	case unicode.IsLetter(char): return LETTER
	case unicode.IsSpace(char): return WHITESPACE
	}
	switch char {
	case '=': return ASSIGN_OP
	case '+': return ADD_OP
	case '-': return SUB_OP
	case '*': return MULT_OP
	case '/': return DIV_OP
	case '(': return LEFT_PAREN
	case ')': return RIGHT_PAREN
	case ';': return EXPR_END
	case '{': return LEFT_BRAC
	case '}': return RIGHT_BRAC
	case '>': return GREAT_OP
	case '<': return LESS_OP
	}
	return UNKNOWN
}
