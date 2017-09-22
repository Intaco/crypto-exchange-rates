package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"errors"
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
	_db.AutoMigrate(&Price{})
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

func RetrieveCurrencies(duration string) (CurrenciesArr, error) {
	var results []Price
	var currs CurrenciesArr
	switch duration {
	case "DAY":
		_db.Select("id, avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment ").
		Where("date(moment) >= date(current_date, '-1 day') AND date(moment) <= current_date").
		Group(`strftime("%H", moment), type`).Order("type, moment").Find(&results)
		break
	case "WEEK":
		_db.Select("id, avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment ").
		Where("date(moment) >= date(current_date, '-7 day') AND date(moment) <= current_date").
		Group(`strftime("%D", moment), type`).Order("type, moment").Find(&results)
		break
	default:
		_db.Select("id, avg(usd) as usd, avg(eur) as eur, avg(rur) as rur, type, moment ").
		Where("date(moment) >= date(current_date, '-30 day') AND date(moment) <= current_date").
		Group(`strftime("%D", moment), type`).Order("type, moment").Find(&results)
		break
	}
	if len(_db.GetErrors()) > 0 {
		log.Printf("Query exec failed. Error: %v", _db.GetErrors())
		return currs, _db.Error
	}
	return splitResultsIntoCurrenciesArr(results)
}
const CURR_COUNT = 9

func splitResultsIntoCurrenciesArr(prices[] Price) (CurrenciesArr, error) {
	var currs CurrenciesArr
	if len(prices) == 0 || len(prices) % CURR_COUNT != 0 {
		return currs, errors.New("incorrect prices rows count for splitting")
	}
	arrLen := len(prices) / CURR_COUNT
	currs.BTC = prices[0:arrLen]
	currs.ETH = prices[arrLen:arrLen*2]
	currs.LTC = prices[arrLen*2:arrLen*2]
	currs.XMR = prices[arrLen*3:arrLen*4]
	currs.ETC = prices[arrLen*4:arrLen*5]
	currs.DASH = prices[arrLen*5:arrLen*6]
	currs.MAID = prices[arrLen*6:arrLen*7]
	currs.REP = prices[arrLen*7:arrLen*8]
	currs.XEM = prices[arrLen*8:arrLen*CURR_COUNT]
	return currs, nil
}