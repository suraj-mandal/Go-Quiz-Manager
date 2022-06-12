// Package utils provides utilities to perform
// 1. Shuffling of the list of Questions
package utils

import (
	"Go_Quiz_Manager/models"
	"math/rand"
	"time"
)

// Question model
type Question = models.Question

// ShuffleQuestions - this function shuffles the list of questions. Basically
// randomizes the indexes of the list and then creates a new list
// with randomized indexes
func ShuffleQuestions(questionList []Question) []Question {

	// setting a random seed here
	rand.Seed(time.Now().UnixNano())

	totalQuestions := len(questionList)
	shuffledList := make([]Question, totalQuestions)
	// randomizing the indexes
	shuffledIndexes := rand.Perm(totalQuestions)
	// shuffling through the list
	for idx, shuffledIdx := range shuffledIndexes {
		shuffledList[shuffledIdx] = questionList[idx]
	}
	// returning the shuffled list
	return shuffledList
}
