package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {

	words := []string{"golang", "java", "php", "js"}
	word := words[rand.Intn(len(words))] // get a random word from the list
	lives := 2 * len(words)              // number of lives

	blanks := []string{}
	for range word {
		blanks = append(blanks, "_")
	}

	for {
		fmt.Printf("💚 %d, Word: %s Letter: ", lives, strings.Join(blanks, " "))
		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		// Check if the user has won or lost
		for _, inputeLetter := range input {
			correctGuess := false
			for i, wordLetter := range word {
				if inputeLetter == wordLetter {
					blanks[i] = string(inputeLetter)
					correctGuess = true
				}
			}

			if !correctGuess {
				lives--
			}
		}

		if lives <= 0 {
			fmt.Printf("💚 0, Word: %s - sorry, You lost!\n", word)
			break
		}
		if word == strings.Join(blanks, "") {
			fmt.Printf("❤️ %d, Word: %s - Congratulations, You won!\n", lives, word)
			break
		}
	}

}
