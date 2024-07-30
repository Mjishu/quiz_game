package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Help function")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println()
	for _, cmd := range get_commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

type question struct {
	question string
	answer   []string
	correct  string
}

func commandStart() error {
	game_answers := []question{
		{
			question: "How many shadows do you have",
			answer:   []string{"1", "2", "3", "4"},
			correct:  "1",
		},
		{
			question: "what is 5 * 5",
			answer:   []string{"5", "10", "25", "125"},
			correct:  "25", //change this to be 3 if i want it to be the index instead?
		},
	}
	game_scanner := bufio.NewScanner(os.Stdin)
	score := 0
	for i := 0; i < len(game_answers); i++ {
		fmt.Println()
		fmt.Println(game_answers[i].question)
		fmt.Println(game_answers[i].answer)
		game_scanner.Scan()
		input := game_scanner.Text()
		if input == game_answers[i].correct {
			fmt.Printf("Correct! The answer was %v\n", input)
			score += 1
		} else {
			fmt.Printf("Incorrect! you entered %v but the answer was %v\n", input, game_answers[i].correct)
		}
		fmt.Println()
	}
	fmt.Println()
	if score > 1 {
		fmt.Printf("Wow you're good! you got %v/%v correct\n", score, len(game_answers))
	} else {
		fmt.Printf("You didnt do so good... Wanna try again? %v/%v\n", score, len(game_answers))
	}
	fmt.Println()
	return nil
}
