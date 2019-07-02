package model

import (
	xtime "mall/lib/time"
)

type EweiShopGroupsBuyer struct {
	Openid  string
	Avatar  string
	Name    string
	Payedat xtime.Time
	Price   float64
}
