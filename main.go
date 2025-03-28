package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	fileContent, err := os.ReadFile("words.json")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	var words []string
	err = json.Unmarshal(fileContent, &words)
	if err != nil {
		fmt.Println("Error unmarshalling file")
		return
	}

	if len(words) == 0 {
		fmt.Println("No words found")
		return
	}

	// انتخاب تصادفی کلمه از لیست

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
