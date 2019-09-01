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
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS time_series_data (id UNSIGNED BIG INT PRIMARY KEY, instrument INTEGER, value REAL, timestamp UNSIGNED BIG INT)")
    statement.Exec()
    return &MikapodStorage{
        database: database,
    }
}

func (s *MikapodStorage) InsertTimeSeriesData(instrument int32, value float32, t *timestamp.Timestamp) {
    log.Printf("Instrument: %v", instrument)
	log.Printf("Value: %v", value)
	log.Printf("Timestamp: %v", t.Seconds)
    statement, _ := s.database.Prepare("INSERT INTO time_series_data (instrument, value, timestamp) VALUES (?, ?, ?)")
    statement.Exec(instrument, value, t.Seconds)
}
