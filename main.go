package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	endpoint := 0
	startingpoint := 0
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Do you want to specify a range of numbers for this game? y/n = ")
	scanner.Scan()
	if scanner.Text() == "y" {
		fmt.Print("starting point:")
		scanner.Scan()
		startingpoint, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Only use valid numbers")
			return
		}

		fmt.Print("End point of your range:")
		scanner.Scan()
		endpoint, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Only use valid numbers")
			return
		}

		if startingpoint >= endpoint {
			fmt.Println("Starting point should be less than endpoint.")
			return
		}
	} else {
		startingpoint = 1
		endpoint = 100
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomnumber := rng.Intn(endpoint-startingpoint+1) + startingpoint

	fmt.Printf("I have selected a random number between %d and %d\n.", startingpoint, endpoint)
	fmt.Println("Can you guess what it is?")

	attempts := 0

	for {
		fmt.Print("Enter your guess:")
		scanner.Scan()
		input := scanner.Text()

		Guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Your guess should only be valid numbers")
			return
		}

		attempts++

		if Guess > randomnumber {
			fmt.Println("Too High! Try again!")
		} else if Guess < randomnumber {
			fmt.Println("Too low! Try again!")
		} else {
			fmt.Printf("Correct! You've guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
