// Package loader provides the utilities of loading the questions from the file as
// well as utilities regarding shuffling the questions list
package loader

import (
	"Go_Quiz_Manager/models"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Question - importing the Question model
type Question = models.Question

func loadFile(fileName string) io.Reader {
	directoryName := "./csv/"
	file, err := os.OpenFile(directoryName+fileName, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("The file could not be opened")
	}
	return file
}

func LoadQuestions(fileName string) []Question {
	file := loadFile(fileName)
	csvReader := csv.NewReader(file)

	// creating the question list
	var questionList []Question

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Question file is corrupted!")
		}
		questionName, answer := record[0], record[1]
		questionList = append(questionList, Question{Name: questionName, Answer: answer})
	}

	return questionList
}
