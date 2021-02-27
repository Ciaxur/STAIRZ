package main

import (
	"database/sql"
	"time"
)

type EventType string

const (
	EventActive   EventType = "active"
	EventInactive           = "inactive"
	EventUnknown            = "unknown"
)

// CreateRelayDb \
// @param db Database Connection Object
// @param id, name, description Values to Enter into Database
func CreateRelayDb(db *sql.DB, id int64, name string, description string) bool {
	insertDynStmt := `INSERT INTO "Relay" ("id", "name", "description") VALUES($1, $2, $3)`
	_, e := db.Exec(insertDynStmt, id, name, description)
	CheckError(e, "Relay Creation Error")
	return e == nil
}

// CreateEventDb \
// @param db Database Connection Object
// @param id, relayID, eventType, distance Values to Enter into Database
func CreateEventDb(db *sql.DB, id string, relayID int64, eventType EventType, distance float64) bool {
	insertDynStmt := `INSERT INTO "Event" ("ID", "relayID", "type", "distance", "created_at") VALUES($1, $2, $3, $4, $5)`
	_, e := db.Exec(insertDynStmt, id, relayID, eventType, distance, time.Now())
	CheckError(e, "Relay Creation Error")
	return e == nil
}
