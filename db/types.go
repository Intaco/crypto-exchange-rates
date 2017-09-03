package db

import (
	"time"
)

type Currencies struct {
	BTC  BTC
	ETH  ETH
	LTC  LTC
	XMR  XMR
	ETC  ETC
	DASH DASH
	MAID MAID
	REP  REP
	XEM  XEM
}

type Price struct {
	ID     int
	USD    float32   `gorm:"column:USD"`
	EUR    float32   `gorm:"column:EUR"`
	RUR    float32   `gorm:"column:RUR"`
	Moment time.Time `sql:"DEFAULT:current_timestamp"`
}

type BTC struct {
	Price
}

func (BTC) TableName() string {
	return "BTC_prices"
}

type ETH struct {
	Price
}

func (ETH) TableName() string {
	return "ETH_prices"
}

type LTC struct {
	Price
}

func (LTC) TableName() string {
	return "LTC_prices"
}

type XMR struct {
	Price
}

func (XMR) TableName() string {
	return "XMR_prices"
}

type ETC struct {
	Price
}

func (ETC) TableName() string {
	return "ETC_prices"
}

type DASH struct {
	Price
}

func (DASH) TableName() string {
	return "DASH_prices"
}

type MAID struct {
	Price
}

func (MAID) TableName() string {
	return "MAID_prices"
}

type REP struct {
	Price
}

func (REP) TableName() string {
	return "REP_prices"
}

type XEM struct {
	Price
}

func (XEM) TableName() string {
	return "XEM_prices"
}
