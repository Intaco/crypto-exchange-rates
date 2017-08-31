package db

import (
	"time"
)

type Price struct {
	USD    float32  `gorm:"column:USD"`
	EUR    float32  `gorm:"column:EUR"`
	RUR    float32  `gorm:"column:RUR"`
	Moment time.Time `gorm:"column:moment"`
}

type BTC struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

func (BTC) TableName() string {
	return "BTC_prices"
}

type ETH struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type LTC struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type XMR struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type ETC struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type DASH struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type MAID struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type REP struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}

type XEM struct {
	Id    uint `gorm:"primary_key"`
	Price Price
}
