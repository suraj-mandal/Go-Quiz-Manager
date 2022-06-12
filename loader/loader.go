package loader

import (
	"Go_Quiz_Manager/models"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"os"
)

type Question models.Question

func loadFile(fileName string) io.Reader {
	directoryName := "./data/"
	file, err := os.OpenFile(directoryName+fileName, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("The file could not be opened")
	}
	return file
}

func ShuffleQuestions(questionList []Question) []Question {
	totalQuestions := len(questionList)
	shuffledList := make([]Question, totalQuestions)
	shuffledIndexes := rand.Perm(totalQuestions)
	for idx, shuffledIdx := range shuffledIndexes {
		shuffledList[shuffledIdx] = questionList[idx]
	}
	return shuffledList
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
