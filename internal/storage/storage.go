package storage

import (
    "database/sql"
    "time"
    // "fmt"
    // "strconv"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type TimeSeriesDatum struct {
    Id int64
    Instrument int32
    Value float32
    Timestamp time.Time
}

type MikapodDB struct {
    database *sql.DB
}

func InitMikapodDB() (*MikapodDB) {
    // DEVELOPERS NOTE:
    // (1) SQLite3 Fields via https://www.sqlite.org/datatype3.html
    // (2) Learn SQL through W3Schools via https://www.w3schools.com/sql/default.asp
    database, _ := sql.Open("sqlite3", "./mikapod.db")
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS time_series_data (id INTEGER PRIMARY KEY AUTOINCREMENT, instrument INTEGER, value REAL, timestamp UNSIGNED BIG INT)")
    statement.Exec()
    return &MikapodDB{
        database: database,
    }
}

func (s *MikapodDB) InsertTimeSeriesData(instrument int32, value float32, t *time.Time) {
    statement, _ := s.database.Prepare("INSERT INTO time_series_data (instrument, value, timestamp) VALUES (?, ?, ?)")
    statement.Exec(instrument, value, t)
	log.Printf("Executed Insertion")
}

func (s *MikapodDB) ListTimeSeriesData() ([]TimeSeriesDatum){
    rows, _ := s.database.Query("SELECT id, instrument, value, timestamp FROM time_series_data LIMIT 250")
    arr := make([]TimeSeriesDatum, 1)

    var id int64
    var instrument int32
    var value float32
    var timestamp time.Time
    for rows.Next() {
        rows.Scan(&id, &instrument, &value, &timestamp)
        // log.Printf("Rows: %v", rows)
        arr = append(arr, TimeSeriesDatum{
            Id: id,
            Instrument: instrument,
            Value: value,
            Timestamp: timestamp,
        });
    }
    log.Printf("Executed Listing")
    return arr
}

func (s *MikapodDB) DeleteTimeSeriesData(pks []int64) {
    log.Printf("Deleting time-series data with PKs: %v", pks)
    for _, v := range pks {
        statement, _ := s.database.Prepare("DELETE FROM time_series_data WHERE id=(?)")
        statement.Exec(v)
    }
	log.Printf("Executed Deletions")
}
