package main

import (
	"Go_Quiz_Manager/loader"
	"Go_Quiz_Manager/models"
	"fmt"
	"strings"
)

type Question models.Question

func main() {

	fileName := "test1.csv"

	// load the questions from the file name
	questionsList := loader.LoadQuestions(fileName)

	// shuffle the questions list
	questionsList = loader.ShuffleQuestions(questionsList)

	// for each question in the question list, I print out the question, and then I ask the user
	// about the answer to the question

	// once the user enters the answer, I then match it with the answer given and if they are equal
	// then I increment user score

	// total number of questions in the list
	total := len(questionsList)

	// total number of correct questions answered
	correct := 0

	fmt.Println("Welcome to Go quiz....")

	fmt.Println("Answer the following questions.\n")

	for idx, question := range questionsList {
		fmt.Printf("%d. %s = ", idx+1, question.Name)
		// getting the input from the user
		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			userInput = ""
		}
		// trimming the spaces from the front and behind of the string
		userInput = strings.Trim(userInput, " ")

		// if the user input matches the answer then increment correct by 1
		if strings.EqualFold(userInput, question.Answer) {
			correct++
		}
	}

	// finally getting the results
	percentageCorrect := float32(correct*100.0) / float32(total)
	fmt.Printf("\nYour score: %.2f\n", percentageCorrect)

}
