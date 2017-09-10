package db

import (
	"time"
)

type Currencies struct {
	BTC  Price
	ETH  Price
	LTC  Price
	XMR  Price
	ETC  Price
	DASH Price
	MAID Price
	REP  Price
	XEM  Price
}

func (c * Currencies) InitTypes() { //with default types
	c.BTC.Type = "BTC"
	c.ETH.Type = "ETH"
	c.LTC.Type = "LTC"
	c.XMR.Type = "XMR"
	c.ETC.Type = "ETC"
	c.DASH.Type = "DASH"
	c.MAID.Type = "MAID"
	c.REP.Type = "REP"
	c.XEM.Type = "XEM"
}

type Price struct {
	ID     int
	USD    float32
	EUR    float32
	RUR    float32
	Type   string    `sql:"not null"`
	Moment time.Time `sql:"DEFAULT:current_timestamp"`
}
