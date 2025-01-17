package data

import (
	"fmt"

	"github.com/qdm12/ddns-updater/internal/models"
	"github.com/qdm12/ddns-updater/internal/records"
)

var _ PersistentDatabase = (*database)(nil)

type PersistentDatabase interface {
	GetEvents(domain, host string) (events []models.HistoryEvent, err error)
	Update(id int, record records.Record) (err error)
	Close() (err error)
}

func (db *database) GetEvents(domain, host string) (events []models.HistoryEvent, err error) {
	return db.persistentDB.GetEvents(domain, host)
}

func (db *database) Update(id int, record records.Record) (err error) {
	db.Lock()
	defer db.Unlock()
	if id < 0 {
		return fmt.Errorf("id %d cannot be lower than 0", id)
	}
	if id > len(db.data)-1 {
		return fmt.Errorf("no record config found for id %d", id)
	}
	currentCount := len(db.data[id].History)
	newCount := len(record.History)
	db.data[id] = record
	// new IP address added
	if newCount > currentCount {
		if err := db.persistentDB.StoreNewIP(
			record.Settings.Domain(),
			record.Settings.Host(),
			record.History.GetCurrentIP(),
			record.History.GetSuccessTime(),
		); err != nil {
			return err
		}
	}
	return nil
}

func (db *database) Close() (err error) {
	db.Lock() // ensure write operation finishes
	defer db.Unlock()
	return db.persistentDB.Close()
}
