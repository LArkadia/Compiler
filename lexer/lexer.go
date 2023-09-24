package lexer

import (
	"compiler/lexer/trees"
	"encoding/json"
	"fmt"
	"os"
)

type tokens_agroupated struct {
	Token_type string   `json:"type"`
	Spaced     bool     `json:"spaced"`
	Operators  []string `json:"operators"`
	Operation  []string `json:"operation"`
}
type raw_tokens struct {
	Tokens []tokens_agroupated `json:"tokens"`
}

func load_raw_tokens(json_source string) raw_tokens {
	raw_tokens_file, err := os.ReadFile(json_source)
	var json_tokens raw_tokens
	if err != nil {
		fmt.Println("Error while reading tokens file!!")
	}
	if err := json.Unmarshal(raw_tokens_file, &json_tokens); err != nil {
		fmt.Println("Error while unmarshalling the tokens!!")
	}
	return json_tokens
}
func unwarp_tokens(json_tokens raw_tokens) (*trees.Node, map[rune]bool, map[string]string) {
	tokens_sequences, non_spaced_tokens, tokens := new(trees.Node), make(map[rune]bool), make(map[string]string)
	for _, token_group := range json_tokens.Tokens {
		for i, token_operator := range token_group.Operators {
			tokens[token_operator] = token_group.Token_type + token_group.Operation[i]
			if !token_group.Spaced {
				non_spaced_tokens[[]rune(token_operator)[0]] = true
			}
			tokens_sequences.Add([]rune(token_operator))
		}
	}
	return tokens_sequences, non_spaced_tokens, tokens
}
func Load_lexer_definition(json_source_path string) (*trees.Node, map[rune]bool, map[string]string) {
	return unwarp_tokens(load_raw_tokens(json_source_path))
}
