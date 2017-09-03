package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var _db *gorm.DB

func InitDb(path string) error {
	var err error
	_db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Migrate the schema
	_db.AutoMigrate(&BTC{}, &ETH{}, &LTC{}, &XMR{}, &ETC{}, &DASH{}, &MAID{}, &REP{}, &XEM{})
	if _db.Error != nil {
		log.Printf("Failed to migrate tables! Error: %v\n", err)
		return err
	}
	return err
}

func WriteCurrencies(currencies Currencies) []error {
	_db = _db.Create(&currencies.BTC)
	_db = _db.Create(&currencies.ETH)
	_db = _db.Create(&currencies.LTC)
	_db = _db.Create(&currencies.XMR)
	_db = _db.Create(&currencies.ETC)
	_db = _db.Create(&currencies.DASH)
	_db = _db.Create(&currencies.MAID)
	_db = _db.Create(&currencies.REP)
	_db = _db.Create(&currencies.XEM)
	if len(_db.GetErrors()) > 0 {
		log.Printf("Failed to write currencies to DB! Error: %v\n", _db.GetErrors())
		return _db.GetErrors()
	}
	return nil
}
