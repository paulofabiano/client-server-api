package database

import (
	"context"
	"strconv"

	"github.com/paulofabiano/client-server-api/server/api"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Quotation struct {
	ID  int `gorm:"primaryKey"`
	Bid float64
	gorm.Model
}

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Quotation{})

	return db, nil
}

func SaveQuotation(ctx context.Context, db *gorm.DB, quotation *api.Quotation) error {
	convertedBid, err := strconv.ParseFloat(quotation.USD.Bid, 64)
	if err != nil {
		return err
	}

	q := Quotation{Bid: convertedBid}
	result := db.WithContext(ctx).Create(&q)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
