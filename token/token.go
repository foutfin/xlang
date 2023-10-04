package token

import (
    "fmt"
)


type TokenType string

type Token struct {
    Literal string
    Type TokenType   
}


const (

    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers 
    IDENT = "IDENT"

    // Literals
    INT = "INT" 
    STR = "STR"

    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"
    LT = "<"
    GT = ">"

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
    TRUE = "TRUE"
    FALSE = "FALSE"

    EQ = "=="
    NOTEQ = "!="

)


var keywords = map[string]TokenType{
    "fn" : FUNCTION,
    "let" : LET,
    "if" : IF,
    "else" : ELSE,
    "return" : RETURN,
    "true" : TRUE,
    "false" : FALSE,

}


func FindType(lit string ) TokenType{
    if tok , ok := keywords[lit] ; ok{
        return tok
    }

    return IDENT
}


func Main(){
    fmt.Println("This is token module")

}
