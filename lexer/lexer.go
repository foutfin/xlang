package lexer

import (
    "boolang/token"
)

type Lexer struct {
    
    input string  // Input data
    curCh byte    // this character that gonna be read
    position int  // this is the last read postion
    readPosition int  // this is the postion of the character to be read

}

func New(input string) *Lexer{
    l := &Lexer{input:input}
    l.readChar() // move the character to first position
    return l
}


func (l *Lexer)readChar(){
     if l.readPosition >= len(l.input){
        l.curCh = 0
     }else{
        l.curCh = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
    for l.curCh == ' ' || l.curCh == '\t' || l.curCh == '\n' || l.curCh == '\r' {
        l.readChar()
    }
}

func (l *Lexer) readIdentifierLiteral() string{
    pos := l.position

    for isLetter(l.curCh){
        l.readChar()
    }

    return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string{
    pos := l.position

    for isDigit(l.curCh){
        l.readChar()
    }
    return l.input[pos:l.position]
}

func (l *Lexer) readString(b byte) string{
    l.readChar()
    pos := l.position

    for l.curCh != b{
        l.readChar()
    }
    final := l.position
    return l.input[pos:final]
}

func (l *Lexer) peekChar()byte{
    if l.readPosition >= len(l.input){
        return 0
    }
    return l.input[l.readPosition]
}

func (l *Lexer) NextToken() token.Token{
    tok := token.Token{}
    l.skipWhitespace()
    switch l.curCh{
        case '=' :
            if l.peekChar() == '='{
                l.readChar()
                tok = token.Token{Type:token.EQ , Literal: "=="}
            }else{
                tok = token.Token{Type : token.ASSIGN , Literal:string(l.curCh) }
            }
        case '(' :
            tok = token.Token{Type : token.LPAREN , Literal:string(l.curCh) }
        case ')' :
            tok = token.Token{Type : token.RPAREN , Literal:string(l.curCh) }
        case '{' :
            tok = token.Token{Type : token.LBRACE , Literal: string(l.curCh) }
        case '}' :
            tok = token.Token{Type : token.RBRACE , Literal: string(l.curCh) }
        case ';' :
            tok = token.Token{Type : token.SEMICOLON , Literal: string(l.curCh) }
        case '+' :
            tok = token.Token{Type : token.PLUS , Literal: string(l.curCh) }
        case ',' :
            tok = token.Token{Type : token.COMMA , Literal: string(l.curCh) }
        case '!' :
            if l.peekChar() == '=' {
                l.readChar()
                tok = token.Token{Type : token.NOTEQ , Literal : "!="}
            }else{
                tok = token.Token{Type : token.BANG , Literal: string(l.curCh) }
            }
        case '-' :
            tok = token.Token{Type : token.MINUS , Literal: string(l.curCh) }
        case '*' :
            tok = token.Token{Type : token.ASTERISK , Literal: string(l.curCh) }
        case '/' :
            tok = token.Token{Type : token.SLASH , Literal: string(l.curCh) }
        case '<' :
            tok = token.Token{Type : token.LT , Literal: string(l.curCh) }
        case '>' :
            tok = token.Token{Type : token.GT , Literal: string(l.curCh) }
        case '"' :
            tok = token.Token{Type : token.STR , Literal: l.readString('"') }
        case '\'' :
            tok = token.Token{Type : token.STR , Literal: l.readString('\'') }

        case 0 :
            tok = token.Token{Type : token.EOF , Literal: "" }

        // Now to identify if it is keyword or identifier or literal because they can contain multiple characters so case will
        // not match up , a default case is there to match those cases

        default :
           if isLetter(l.curCh) {
                tok.Literal = l.readIdentifierLiteral()
                tok.Type = token.FindType(tok.Literal)
                return tok
           } else if isDigit(l.curCh){
               tok = token.Token{Type: token.INT , Literal : l.readNumber()}
               return tok
           }else{
               tok = token.Token{Type: token.ILLEGAL,Literal: string(l.curCh)} 
           }
    }

    l.readChar()
    return tok

}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool{
    return '0' <= ch && ch <= '9'
}

