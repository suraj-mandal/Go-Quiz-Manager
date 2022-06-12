package main

import (
	"Go_Quiz_Manager/loader"
	"Go_Quiz_Manager/models"
	"Go_Quiz_Manager/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Question - creating type alias here
type Question = models.Question

func getUserInputChannel(inputChannel chan string, reader bufio.Reader) {

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// trimming carriage return and newline if there are any
	inputChannel <- strings.Trim(name, "\n\r")

}

func main() {

	fileName := "test1.csv"

	reader := bufio.NewReader(os.Stdin)

	// load the questions from the file name
	questionsList := loader.LoadQuestions(fileName)

	questionsList = utils.ShuffleQuestions(questionsList)

	// for each question in the question list, I print out the question, and then I ask the user
	// about the answer to the question

	// once the user enters the answer, I then match it with the answer given and if they are equal
	// then I increment user score

	// total number of questions in the list
	total := len(questionsList)

	// total number of correct questions answered
	correct := 0
	attempted := 0

	fmt.Println("\nWelcome to Go quiz....")

	fmt.Println("Answer the following questions.")

	fmt.Println()

	c := make(chan string)

	for idx, question := range questionsList {
		fmt.Printf("%d. %s = ", idx+1, question.Name)

		// getting the input from the user
		go getUserInputChannel(c, *reader)

		select {
		case userInput := <-c:
			// if the user input matches the answer then increment correct by 1

			if strings.EqualFold(userInput, question.Answer) {
				correct++
			}

			attempted++

		case <-time.After(3010 * time.Millisecond):
			fmt.Println()
		}

	}

	percentageCorrect := float32((correct)*100.0) / float32(total)
	fmt.Printf("\n\nTotal Questions attempted: %d", attempted)
	fmt.Printf("\nTotal Questions correct: %d", correct)
	fmt.Printf("\nYour score: %.2f\n", percentageCorrect)

}
