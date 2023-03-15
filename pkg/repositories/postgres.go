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
	if err := migration(cfg); err != nil {
		return nil, err
	}

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migration(cfg Config) error {
	dbc, err := sqlx.Connect("postgres", "host="+cfg.Host+" user="+cfg.Username+" dbname="+cfg.DBName+" port="+cfg.Port+" sslmode="+cfg.SSLMode+" password="+cfg.Password)
	if err != nil {
		logrus.Printf("migration : Register DB : %v", err)
	}
	defer dbc.Close()

	m := sqlxmigrate.New(dbc, sqlxmigrate.DefaultOptions, []*sqlxmigrate.Migration{
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
