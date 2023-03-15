package models

type Users struct {
	Id              int64
	ChatID          int64
	Username        string
	FirstQueryTime  string
	AllQueriesCount int
}

func NewUsers(chatID int64, username string, firstQueryTime string, allQueriesCount int) *Users {
	return &Users{ChatID: chatID, Username: username, FirstQueryTime: firstQueryTime, AllQueriesCount: allQueriesCount}
}
