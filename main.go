package main

import (
	"github.com/intaco/crypto-exchange-rates/db"
	"github.com/intaco/crypto-exchange-rates/updater"
)

const DBPATH = "crypto-exchange.db"

func main() {
	messages := make(chan int)
	db.InitDb(DBPATH)
	go updater.InitUpdater()
	<-messages
}
