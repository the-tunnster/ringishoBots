package models

type Question struct {
	UserID        string `json:"UserID"`
	QuestionID    string `json:"QuestionID"`
	BotID         int    `json:"BotID"`
	BotParameters string `json:"BotParams"`
}

/*
{
	"UserID" : "123",
	"QuestionID" : "456",
	"BotID" : 789,
	"BotParams" : "Pune"
}
*/

type Answer struct {
	UserID     string `json:"UserID"`
	QuestionID string `json:"QuestionID"`
	Status     string `json:"Status"`
}

type Bot struct {
	BotID   int    `json:"BotID"`
	BotName string `json:"BotName"`
}

type AlgoliaRecord struct{
	ObjectID string `json:"objectID"`
	Name string `json:"name"`
}