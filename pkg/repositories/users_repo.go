package repositories

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"telegram_bot_golang/pkg/models"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *models.Users) error {
	var query string
	var row *sql.Row
	if user.AllQueriesCount == 0 {
		query = fmt.Sprintf("INSERT INTO %s (chat_id, username, first_query_time, all_queries_count) values ($1, $2, $3, $4) RETURNING id", "users")

		row = r.db.QueryRow(query, user.ChatID, user.Username, user.FirstQueryTime, 0)
	} else {
		query = fmt.Sprintf("INSERT INTO %s (chat_id, username, first_query_time, all_queries_count) values ($1, $2, $3, $4) RETURNING id", "users")

		row = r.db.QueryRow(query, user.ChatID, user.Username, user.FirstQueryTime, user.AllQueriesCount)
	}

	var id int
	if err := row.Scan(&id); err != nil {
		logrus.Printf("Create User failed %s", err.Error())
		return err
	}
	logrus.Print("Insert into Users")

	return nil
}

func (r *UserRepo) UpdateUserAllQueriesCount(charID int64) error {

	query := fmt.Sprintf(`UPDATE %s u SET all_queries_count = all_queries_count + 1
									WHERE u.chat_id = $1`,
		"users")

	_, err := r.db.Exec(query, charID)
	if err != nil {
		logrus.Printf("Failed Update User All Queries Count: %s", err.Error())
		return err
	}

	logrus.Print("Update User All Queries Count")

	return nil
}

func (r *UserRepo) GetUserFirstQueryTime(userID int64) (string, error) {
	query := fmt.Sprintf(`SELECT u.first_query_time FROM %s u 
                     				WHERE u.id = $1 `,
		"users")
	var description string
	if err := r.db.Get(&description, query, userID); err != nil {
		logrus.Printf("Get User First Query Time failed %s", err.Error())
		return description, err
	}

	logrus.Print("Get User First Query Time")

	return description, nil
}

func (r *UserRepo) GetUserAllQueriesCount(chatID int64) (int, error) {
	query := fmt.Sprintf(`SELECT u.all_queries_count FROM %s u 
                     				WHERE u.chat_id= $1 `,
		"users")

	var AllQueriesCount int
	if err := r.db.Get(&AllQueriesCount, query, chatID); err != nil {
		logrus.Printf("GetUserAllQueriesCount failed %s", err.Error())
		return 0, err
	}

	logrus.Printf("Get User All Queries Count %d", AllQueriesCount)

	return AllQueriesCount, nil
}

func (r *UserRepo) GetUserID(chatID int64) (int, error) {
	query := fmt.Sprintf(`SELECT u.id FROM %s u 
                     				WHERE u.chat_id = $1 `,
		"users")
	var AllQueriesCount int
	if err := r.db.Get(&AllQueriesCount, query, chatID); err != nil {
		logrus.Printf("GetUserAllQueriesCount failed %s", err.Error())
		return AllQueriesCount, err
	}

	logrus.Print("Get User ID")

	return AllQueriesCount, nil
}

func (r *UserRepo) CheckUserByID(userID int64) string {
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);`)
	var check string
	if err := r.db.Get(&check, query, userID); err != nil {
		logrus.Printf("CheckUserByID failed %s", err.Error())
	}

	logrus.Print("Get User ID")

	return check
}
