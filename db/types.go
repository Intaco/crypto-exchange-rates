package db

import (
	"time"
)

type Price struct {
	ID  int
	USD float32  `gorm:"column:USD"`
	EUR float32  `gorm:"column:EUR"`
	RUR float32  `gorm:"column:RUR"`
}

type BTC struct {
	ID      int
	Price   Price
	PriceID int
	Moment  time.Time
}

func (BTC) TableName() string {
	return "BTC_prices"
}

type ETH struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (ETH) TableName() string {
	return "ETH_prices"
}

type LTC struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (LTC) TableName() string {
	return "LTC_prices"
}

type XMR struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (XMR) TableName() string {
	return "XMR_prices"
}

type ETC struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (ETC) TableName() string {
	return "ETC_prices"
}

type DASH struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (DASH) TableName() string {
	return "DASH_prices"
}

type MAID struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (MAID) TableName() string {
	return "MAID_prices"
}

type REP struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (REP) TableName() string {
	return "REP_prices"
}

type XEM struct {
	ID      uint
	Price   Price
	PriceID int
	Moment  time.Time
}

func (XEM) TableName() string {
	return "XEM_prices"
}
