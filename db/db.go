package db

import (
	_ "github.com/mattn/go-sqlite3"
	//"sync"
	"log"
	//"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	defer db.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}
	/*db = db.Exec(DB_INITILIZE_SQL_31_08_2017)*/
	// Migrate the schema
	db.AutoMigrate(&BTC{}, &Price{}, &ETH{}, &LTC{}, &XMR{}, &ETC{}, &DASH{}, &MAID{}, &REP{}, &XEM{})
	if db.Error != nil {
		log.Printf("Failed to migrate tables! Error: %s\n", err)
		return err
	}
	return err
}
