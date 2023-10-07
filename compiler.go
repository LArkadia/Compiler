package main

import (
	"compiler/lexer"
	"fmt"
	"os"
)

func main() {
	params := os.Args[1:]
	switch {
	case len(params) == 0:
		fmt.Println("Please select a code file or the -CLI flag")
	case len(params) == 1:
		switch params[0] {
		case "-CLI":

		default:
			lexed_code := lexer.Lexate_file(params[0])
			fmt.Println(lexed_code)
		}
	}
	//NTree := trees.Get_numbers_tree()
	//WS := trees.Get_white_spaces()
	//tokens_sequences, non_spaced_tokens, tokens := lexer.Load_lexer_definition("tokens.json")
	//lexed_code := lexer.Lexate([]rune("var num12 = 3.5E3.12"), tokens_sequences, non_spaced_tokens, tokens, NTree, WS)
	//fmt.Println(lexed_code)
}
