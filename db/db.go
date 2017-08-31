package db

import (
	_ "github.com/mattn/go-sqlite3"
	//"sync"
	"log"
	//"time"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Temp(){
	//Test()
}

func InitDb(path string) error {
	var err error
	db, err = gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	db = db.Exec(DB_INITILIZE_SQL_31_08_2017)
	db.Create(&BTC{Price : Price{USD:0.0, EUR:0.0, RUR: 0.0, Moment: time.Now()}})
	if db.Error != nil {
		log.Printf("%q: %s\n", err, DB_INITILIZE_SQL_31_08_2017)
		return err
	}
	return err
}

const DB_INITILIZE_SQL_31_08_2017 string = `
	CREATE TABLE if not exists BTC_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists ETH_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists LTC_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists XMR_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists ETC_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists DASH_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists MAID_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists REP_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
	CREATE TABLE if not exists XEM_prices (id INTEGER PRIMARY KEY AUTOINCREMENT,
		USD REAL, EUR REAL, RUR REAL, moment DATETIME NOT NULL);
`

/*
func Test() error {
	mu.Lock()
	defer mu.Unlock()
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln("Failed to connect to db! Error: %s", err)
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO BTC_prices (USD, EUR, RUR, moment) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatalln("Failed to prepare statement! Error: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(0.0, 0.0, 0.0, time.Now())
	if err != nil {
		log.Fatalln("Failed to exec statement! Error: %s", err)
		return err
	}
	tx.Commit()
	return nil
}*/
