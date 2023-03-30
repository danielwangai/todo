package psql

import (
	"database/sql"
	"fmt"

	"github.com/danielwangai/todo-app/internal/svc"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type dbClient struct {
	db  *sql.DB
	log *logrus.Logger
}

func New(db *sql.DB, log *logrus.Logger) svc.DAO {
	return &dbClient{db, log}
}

// NewDBClient creates a new database instance
func NewDBClient(log *logrus.Logger, username, password, host, dbName, sslMode string) (*sql.DB, error) {
	// connString := fmt.Sprintf("postgresql://%s:%s@%s/%s??sslmode=disable", username, password, host, dbName)
	connString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=%s", host, username, dbName, password, sslMode)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.WithError(err).Error("Failed to connect to database")
		return nil, err
	}

	return db, nil
}

// PingDB checks if there's a successful connection to the database
func PingDB(log *logrus.Logger, db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.WithError(err).Error("failed to ping databse")
		return err
	}

	log.Infof("successfully connected to database!")
	return nil
}
