package main

import (
	"github.com/intaco/crypto-exchange-rates/db"
	"github.com/intaco/crypto-exchange-rates/updater"
	"github.com/intaco/crypto-exchange-rates/ui"
	//"time"
)

const DBPATH = "crypto-exchange.db"

func main() {
	messages := make(chan int)
	db.InitDb(DBPATH)
	go updater.InitUpdater()
	go ui.Serve()
	<-messages
}
