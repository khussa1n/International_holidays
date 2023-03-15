package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"telegram_bot_golang/pkg/models"
)

type DateRepo struct {
	db *sqlx.DB
}

func NewDateRepo(db *sqlx.DB) *DateRepo {
	return &DateRepo{db: db}
}

func (r *DateRepo) CreateDate(date *models.Dates) error {
	query := fmt.Sprintf("INSERT INTO %s (chat_id, description, date) values ($1, $2, $3) RETURNING id", "dates")

	row := r.db.QueryRow(query, date.ChatID, date.Description, date.Date)

	var id int
	if err := row.Scan(&id); err != nil {
		logrus.Printf("Create Date failed %s", err.Error())
		return err
	}

	logrus.Print("Insert into Dates")

	return nil
}

func (r *DateRepo) DeleteDate(userID string, description string) error {
	query := fmt.Sprintf(`DELETE FROM %s d USING %s u
									WHERE d.user_id = $1 AND d.user_id = u.id AND d.id = $2`,
		"dates", "users")
	_, err := r.db.Exec(query, userID, description)

	if err != nil {
		return err
	}

	logrus.Print("Delete Date")

	return nil
}

func (r *DateRepo) GetDateByDate(date string) ([]string, error) {
	query := fmt.Sprintf(`SELECT d.description FROM %s d 
                     				WHERE d.date = $1 `,
		"dates")
	var description []string
	if err := r.db.Select(&description, query, date); err != nil {
		logrus.Printf("GetDateByDate failed %s", err.Error())
		return description, err
	}

	logrus.Print("Get Date By Date")

	return description, nil
}

func (r *DateRepo) GetAllDate(userID int64) ([]models.Dates, error) {
	var dates []models.Dates
	query := fmt.Sprintf(`SELECT * FROM %s d INNER JOIN %s u on u.id = d.user_id
									WHERE u.id = $1`,
		"dates")
	if err := r.db.Select(&dates, query, userID); err != nil {
		logrus.Printf("GetAllDate failed %s", err.Error())
		return nil, err
	}

	logrus.Print("Get All Date")

	return dates, nil
}
