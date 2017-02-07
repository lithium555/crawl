package main

import "fmt"

func main() {
	var guessColor string
	const favColor = "blue"
	for {
		fmt.Println("Guess my favorite color:")
		if _, err := fmt.Scanf("%s", &guessColor); err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		if favColor == guessColor {
			fmt.Printf("%q is my favorite color!", favColor)
			return
		}
		fmt.Printf("Sorry, %q is not my favorite color.  Guess again.\n", guessColor)
	}
}
