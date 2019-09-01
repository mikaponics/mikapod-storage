package storage

import (
    "database/sql"
    // "time"
    // "fmt"
    // "strconv"
    "log"

    _ "github.com/mattn/go-sqlite3"
    "github.com/golang/protobuf/ptypes/timestamp"
)


type MikapodStorage struct {
    database *sql.DB
}

func InitMikapodStorage() (*MikapodStorage) {
    database, _ := sql.Open("sqlite3", "./mikapod.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    statement.Exec()
    return &MikapodStorage{
        database: database,
    }
}

func (s *MikapodStorage) InsertTimeSeriesData(instrument string, value float32, t *timestamp.Timestamp) {
    log.Printf("Instrument: %v", instrument)
	log.Printf("Value: %v", value)
	log.Printf("Timestamp: %v", t)
    // statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
    // statement.Exec("Nic", "Raboy")
}
