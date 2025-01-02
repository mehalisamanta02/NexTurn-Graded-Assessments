package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Question string
	Options  [4]string
	Answer   int
}

var questions = []Question{
	{
		Question: "What is the capital of India?",
		Options:  [4]string{"1. Delhi", "2. Mumbai", "3. Lucknow", "4. Hyderabad"},
		Answer:   1,
	},
	{
		Question: "What is 5 + 2?",
		Options:  [4]string{"1. 6", "2. 9", "3. 7", "4. 2"},
		Answer:   3,
	},
	{
		Question: "What holiday is Santa Claus associated with?",
		Options:  [4]string{"1. Holi", "2. New Year", "3. Christmas", "4. Diwali"},
		Answer:   3,
	},
	{
		Question: "How many hours there in a day?",
		Options:  [4]string{"1. 32", "2. 24", "3. 20", "4. 22"},
		Answer:   2,
	},
	{
		Question: "Rainbow consists of how many colours?",
		Options:  [4]string{"1. 3", "2. 5", "3. 10", "4. 7"},
		Answer:   4,
	},
}

func takeQuiz() int {
	score := 0

	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.Question)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		answerCh := make(chan string)

		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Enter your answer (1-4) or type 'exit' to quit: ")
			if scanner.Scan() {
				answerCh <- scanner.Text()
			}
		}()

		timer := time.NewTimer(5 * time.Second)
		select {
		case input := <-answerCh:
			timer.Stop()
			if strings.ToLower(input) == "exit" {
				fmt.Println("Exiting quiz early.")
				return score
			}

			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}

			if answer == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong answer.")
			}

		case <-timer.C:
			fmt.Println("Time out! Moving to the next question.")
		}
	}

	return score
}

func classifyPerformance(score int, total int) {
	percentage := float64(score) / float64(total) * 100
	fmt.Printf("\nYour Score: %d/%d (%.2f%%)\n", score, total, percentage)

	switch {
	case percentage >= 80:
		fmt.Println("Performance: Excellent")
	case percentage >= 50:
		fmt.Println("Performance: Good")
	default:
		fmt.Println("Performance: Needs Improvement")
	}
}

func main() {
	fmt.Println("Welcome to the Online Examination System")
	score := takeQuiz()
	classifyPerformance(score, len(questions))
}
