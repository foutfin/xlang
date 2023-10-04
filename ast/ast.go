package ast

import(
    "boolang/token"
    "fmt"
)


type Node interface{
    TokenLiteral() string
}

type Statement interface{
    Node
    statementNode()
}

type Expression interface{
    Node
    expressionNode()
}


type Program struct{
    Statements []Statement
}


func (p *Program) TokenLiteral() string{
    if len(p.Statements) < 1{
        return ""
    }

    return p.Statements[0].TokenLiteral()
}

func (p *Program) String() {
    for _ , s := range p.Statements{
        fmt.Printf("Token is %d \n",s.TokenLiteral())
    }
}


type LetStatement struct{
    Token token.Token
    Name Identifier
    Value Expression
}

func (ls *LetStatement) statementNode(){}

func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

type Identifier struct {
    Token token.Token
    Value string
}


func (i *Identifier) expressionNode(){}

func (i *Identifier) TokenLiteral() string{
    return i.Token.Literal
}




