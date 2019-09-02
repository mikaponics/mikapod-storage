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


type MikapodDB struct {
    database *sql.DB
}

func InitMikapodDB() (*MikapodDB) {
    // DEVELOPERS NOTE:
    // (1) SQLite3 Fields via https://www.sqlite.org/datatype3.html
    // (2) Learn SQL through W3Schools via https://www.w3schools.com/sql/default.asp
    database, _ := sql.Open("sqlite3", "./mikapod.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS time_series_data (id UNSIGNED BIG INT PRIMARY KEY, instrument INTEGER, value REAL, timestamp UNSIGNED BIG INT)")
    statement.Exec()
    return &MikapodDB{
        database: database,
    }
}

func (s *MikapodDB) InsertTimeSeriesData(instrument int32, value float32, t *timestamp.Timestamp) {
    log.Printf("Instrument: %v", instrument)
	log.Printf("Value: %v", value)
	log.Printf("Timestamp: %v", t.Seconds)
    statement, _ := s.database.Prepare("INSERT INTO time_series_data (instrument, value, timestamp) VALUES (?, ?, ?)")
    statement.Exec(instrument, value, t.Seconds)
}
