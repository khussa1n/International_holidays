package repositories

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kulado/sqlxmigrate"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.Printf("Failed Open DB: %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Printf("Failed Ping to DB: %s", err.Error())
		return nil, err
	}

	if err = migration(db); err != nil {
		logrus.Printf("Migration failed: %s", err.Error())
		return nil, err
	}

	return db, nil
}

func migration(db *sqlx.DB) error {
	var err error
	m := sqlxmigrate.New(db, sqlxmigrate.DefaultOptions, []*sqlxmigrate.Migration{
		// create tables
		{
			ID: "201608301400",
			Migrate: func(tx *sql.Tx) error {
				q := `CREATE TABLE users (
						id                serial primary key,
						chat_id           integer unique not null,
						username          varchar(255) not null,
						first_query_time  varchar(255),
						all_queries_count integer
					);`
				_, err = tx.Exec(q)
				return err
			},
			Rollback: func(tx *sql.Tx) error {
				q := `DROP TABLE users;`
				_, err = tx.Exec(q)
				return err
			},
		},
		{
			ID: "201608301415",
			Migrate: func(tx *sql.Tx) error {
				q := `CREATE TABLE dates (
					   id                serial primary key,
					   chat_id           integer references users(chat_id),
					   description       text not null,
					   date          varchar(30)
					);`
				_, err = tx.Exec(q)
				return err
			},
			Rollback: func(tx *sql.Tx) error {
				q := `DROP TABLE dates;`
				_, err = tx.Exec(q)
				return err
			},
		},
	})

	if err = m.Migrate(); err != nil {
		logrus.Printf("Could not migrate: %v", err)
		return err
	}

	logrus.Printf("Migration did run successfully")
	return nil
}
