package main

import (
	"github.com/intaco/crypto-exchange-rates/updater"
	"github.com/intaco/crypto-exchange-rates/db"
)

const DBPATH = "crypto-exchange.db"

func main() {
	messages := make(chan int)
	go updater.InitUpdater()
	db.InitDb(DBPATH)
	<- messages
}
