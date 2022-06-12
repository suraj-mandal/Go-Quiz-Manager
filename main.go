package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Question struct {
	name   string
	answer string
}

func main() {

	// reading the csv file and storing it in the io.reader object along with the error
	f, err := os.OpenFile("./data/test1.csv", os.O_RDONLY, 0755)
	// if error is not null then file could not be opened
	if err != nil {
		log.Fatal("The file could not be opened in read mode")
	}

	// creating a new csv reader to read the file contents
	reader := csv.NewReader(f)

	// will store the list of questions
	var questionsList []Question

	// iterating through each row in the reader object
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error parsing the file")
		}

		question, answer := record[0], record[1]

		questionsList = append(questionsList, Question{name: question, answer: answer})

	}

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
		fmt.Printf("%d. %s = ", idx+1, question.name)
		// getting the input from the user
		var userInput string
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			userInput = ""
		}
		// trimming the spaces from the front and behind of the string
		userInput = strings.Trim(userInput, " ")

		// if the user input matches the answer then increment correct by 1
		if strings.EqualFold(userInput, question.answer) {
			correct++
		}
	}

	// finally getting the results
	percentageCorrect := float32(correct*100.0) / float32(total)
	fmt.Printf("Your score: %.2f\n", percentageCorrect)

}
