package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var timelineDb *sql.DB = nil

func InsertTimelineEvent(account string, block int64, when time.Time, amount int64) error {
	statement, err := timelineDb.Prepare("INSERT INTO timeline (id, account, block, time, amount) VALUES (null, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("unable to prepare timeline insert query: %v", err)
		return err
	}
	_, err = statement.Exec(account, block, when, amount)
	if err != nil {
		log.Printf("unable to exec timeline insert query: %v", err)
		return err
	}
	return nil
}

func init() {
	var err error

	timelineDb, err = sql.Open("sqlite3", "/home/app/.pocket/data/timeline.db?_journal_mode=WAL")
	if err != nil {
		log.Fatalf("unable to open timeline db: %v", err)
	}
	statement, err := timelineDb.Prepare("CREATE TABLE IF NOT EXISTS timeline(id INTEGER NOT NULL PRIMARY KEY, account TEXT, block INTEGER, time DATETIME, amount INTEGER)")
	if err != nil {
		log.Fatalf("unable to prepare create table statement in timeline db: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("unable to create table in timeline db: %v", err)
	}
}
