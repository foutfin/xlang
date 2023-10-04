package parser

import(
    "boolang/ast"
    "boolang/lexer"
    "boolang/token"
    "fmt"
)


type Parser struct {
    l *lexer.Lexer
    curToken token.Token
    peekToken token.Token
    errors []string
}

func (p *Parser) Errors() []string{
    return p.errors
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s instead",t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    }
    p.peekError(t)
    return false
}

func (p *Parser)nextToken(){
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser)ParseProgram() *ast.Program{
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for !p.curTokenIs(token.EOF) {
        statement := p.ParseStatement()
        if statement != nil{
            program.Statements = append(program.Statements,statement)
        }
        p.nextToken()
    }
    return program
}

func (p *Parser) ParseStatement() ast.Statement{
    switch p.curToken.Type {
        case token.LET :
            return p.ParseLetStatement()
        default :
            return nil
    }
}

func (p *Parser) ParseLetStatement() *ast.LetStatement{
    statement := &ast.LetStatement{ Token : p.curToken }

    if !p.expectPeek(token.IDENT){
        return nil
    }

    statement.Name = ast.Identifier{Token : p.curToken , Value : p.curToken.Literal}

    if !p.expectPeek(token.ASSIGN){
        return nil
    }

    for !p.curTokenIs(token.SEMICOLON){
        if(p.curTokenIs(token.EOF)){
            msg := fmt.Sprintf("expected ; (semi colon) at the end of the statement")
            p.errors = append(p.errors, msg)
            return nil
        }
        p.nextToken()
    }
    return statement

}

func New(l *lexer.Lexer) *Parser{
    p := &Parser{l:l,errors : []string{}}
    p.nextToken()
    p.nextToken() // calling two times because to set the curToken and peekToken
    return p
}




