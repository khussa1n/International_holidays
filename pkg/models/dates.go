package models

type Dates struct {
	Id          int64
	ChatID      int64
	Description string
	Date        string
}

func NewDates(ChatID int64, description string, date string) *Dates {
	return &Dates{ChatID: ChatID, Description: description, Date: date}
}
