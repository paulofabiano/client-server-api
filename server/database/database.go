package database

import (
	"context"
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Quotation struct {
	ID  int `gorm:"primaryKey"`
	Bid float64
	gorm.Model
}

func InitDatabase() (*sql.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SaveQuotation(ctx context.Context, db *sql.DB, quotation *Quotation) error {
	stmt, err := db.Prepare("INSERT INTO quotations (bid) VALUES (?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(quotation.Bid)
	if err != nil {
		return err
	}

	return nil
}
