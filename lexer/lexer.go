package lexer

import (
	"compiler/lexer/trees"
	"encoding/json"
	"fmt"
	"os"
)

type tokensGrouped struct {
	TokenType string   `json:"type"`
	Spaced    bool     `json:"spaced"`
	Operators []string `json:"operators"`
	Operation []string `json:"operation"`
}
type rawTokens struct {
	Tokens []tokensGrouped `json:"tokens"`
}

func loadRawTokens(jsonSource string) rawTokens {
	rawTokensFile, err := os.ReadFile(jsonSource)
	var jsonTokens rawTokens
	if err != nil {
		fmt.Println("Error while reading tokens file!!")
	}
	if err := json.Unmarshal(rawTokensFile, &jsonTokens); err != nil {
		fmt.Println("Error while unmarshalling the tokens!!")
	}
	return jsonTokens
}
func unwrapTokens(jsonTokens rawTokens) (*trees.Node, map[rune]bool, map[string]string) {
	tokensSequences, nonSpacedTokens, tokens := new(trees.Node), make(map[rune]bool), make(map[string]string)
	for _, tokenGroup := range jsonTokens.Tokens {
		for i, tokenOperator := range tokenGroup.Operators {
			tokens[tokenOperator] = tokenGroup.TokenType + tokenGroup.Operation[i]
			if !tokenGroup.Spaced {
				nonSpacedTokens[[]rune(tokenOperator)[0]] = true
			}
			tokensSequences.Add([]rune(tokenOperator))
		}
	}
	return tokensSequences, nonSpacedTokens, tokens
}
func LoadLexerDefinition(jsonSourcePath string) (*trees.Node, map[rune]bool, map[string]string) {
	return unwrapTokens(loadRawTokens(jsonSourcePath))
}
