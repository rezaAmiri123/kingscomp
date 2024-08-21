package service

import "github.com/rezaAmiri123/kingscomp/steps/10_event/internal/repository"


type QuestionService struct{
	repository.Question
}

func NewQuestionService(rep repository.Question)*QuestionService{
	return &QuestionService{Question: rep}
}