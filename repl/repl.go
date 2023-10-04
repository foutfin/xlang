package repl
import (
    "bufio"
    "fmt"
    "io"
    "boolang/lexer"
    // "boolang/token"
    "boolang/parser"
    )


const PROMPT = ">> "

// func Start(in io.Reader, out io.Writer) {
//     scanner := bufio.NewScanner(in)
//     for {
//         fmt.Printf(PROMPT)
//         scanned := scanner.Scan()
//         if !scanned {
//             return
//         }
//         line := scanner.Text()
//         if line == "quit()" {
//             fmt.Println("\nBANG BANG")
//             return
//         }
//         l := lexer.New(line)
//         noToken := 0
//         for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
//             fmt.Printf("%+v\n", tok)
//             noToken += 1
//         }
//         fmt.Printf("No of Token Found in statement = %d \n",noToken)
//
//     }
// }


func Start(in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    for {
        fmt.Printf(PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }
        line := scanner.Text()
        l := lexer.New(line)
        p := parser.New(l)
        program := p.ParseProgram()
        if len(p.Errors()) != 0 {
            printParserErrors(out, p.Errors())
            continue
        }
        program.String()
        // io.WriteString(out, program.String())
        // io.WriteString(out, "\n")
    }
}

func printParserErrors(out io.Writer, errors []string) {
    for _, msg := range errors {
        io.WriteString(out, "\t"+msg+"\n")
    }
}
