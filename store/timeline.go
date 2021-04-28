package store

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var timelineDb *sql.DB = nil

func InsertTimelineEvent(account string, height int64, amount int64) error {
	statement, err := timelineDb.Prepare("INSERT INTO timeline (account, block, amount) VALUES (?, ?, ?)")
	if err != nil {
		log.Warningf("unable to prepare timeline insert query: %v", err)
		return err
	}
	_, err = statement.Exec(account, height, amount)
	if err != nil {
		log.Warningf("unable to exec timeline insert query: %v", err)
		return err
	}
	return nil
}

func init () {
	timelineDb, err := sql.Open("sqlite3", "timeline.db")
	if err != nil {
		log.Fatalf("unable to open timeline db: %v", err)
	}
	statement, err := timelineDb.Prepare("CREATE TABLE IF NOT EXISTS timeline(account TEXT, block INTEGER, amount INTEGER)")
	if err != nil {
		log.Fatalf("unable to prepare create table statement in timeline db: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("unable to create table in timeline db: %v", err)
	}
}
