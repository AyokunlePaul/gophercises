package main

import (
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
	"fmt"
)

var(
	numberOfCorrectAnswer,
	numberOfWrongAnswer int64
)

func main() {
	var questionSlice []Question
	questionFile, err := os.Open("/home/eipeks/go/src/LearningGo/gophercises/lesson1/problems.csv")
	if err != nil{
		 log.Fatal(err)
	}
	defer questionFile.Close()
	fmt.Println("Loading questions...")
	fileReader := csv.NewReader(bufio.NewReader(questionFile))
	for {
		question, readingError := fileReader.Read()
		if readingError == io.EOF {
			break
		}
		if readingError != nil {
			log.Fatal(readingError)
		}
		currentQuestion := question[0]
		currentAnswer := question[1]
		questionSlice = append(questionSlice, Question{Question:currentQuestion, Answer:currentAnswer})
	}
	fmt.Println("Finished!!!!")

	startAskingQuestion(questionSlice)
	fmt.Printf("%s = %d\n%s = %d",
		"Number of questions you got correctly", numberOfCorrectAnswer,
		"Number of questions you got wrong", numberOfWrongAnswer)
}

func startAskingQuestion(q []Question) {
	for _, currentQuestion := range q{
		fmt.Println(currentQuestion.Question, "= ")
		var userAnswer string
		fmt.Scan(&userAnswer)
		if currentQuestion.Answer == userAnswer {
			fmt.Println("Correct")
			numberOfCorrectAnswer += 1
		} else {
			fmt.Println("Wrong")
			numberOfWrongAnswer += 1
		}
	}
}

type Question struct {
	Question string
	Answer string
}

//func (q *Question) createQuestion(question, answer string){
//	q.Question = question
//	q.Answer = answer
//}
