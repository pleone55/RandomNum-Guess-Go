package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 7

func main() {
	randomNumGuess()
}

func randomNumGuess() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Must provide a guess")
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("YOU WIN!")
			fmt.Println("Number of turns: ", turn)
			return
		}
	}
	fmt.Println("YOU LOSE")
}
