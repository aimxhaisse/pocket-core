package store

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var timelineDb *sql.DB = nil

type TxType = int64
const (
	TxMint TxType = iota
	TxSlash
	TxSimpleSlash
)

func InsertTimelineEvent(account string, height int64, amount int64, txtype TxType) error {
	statement, err := timelineDb.Prepare("INSERT INTO timeline (account, block, amount, txtype) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Printf("unable to prepare timeline insert query: %v", err)
		return err
	}
	_, err = statement.Exec(account, height, amount, txtype)
	if err != nil {
		log.Printf("unable to exec timeline insert query: %v", err)
		return err
	}
	return nil
}

func init () {
	var err error

	timelineDb, err = sql.Open("sqlite3", "timeline.db")
	if err != nil {
		log.Fatalf("unable to open timeline db: %v", err)
	}
	statement, err := timelineDb.Prepare("CREATE TABLE IF NOT EXISTS timeline(account TEXT, block INTEGER, amount INTEGER, txtype INTEGER)")
	if err != nil {
		log.Fatalf("unable to prepare create table statement in timeline db: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("unable to create table in timeline db: %v", err)
	}
}
