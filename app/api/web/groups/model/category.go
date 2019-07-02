package model

type EweiShopGroupsCategory struct {
	Id           int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Uniacid      int    `json:"uniacid"`
	Name         string `json:"name"`
	Thumb        string `json:"thumb"`
	Displayorder int    `json:"displayorder"`
	Enabled      int    `json:"enabled"`
	Advimg       string `json:"advimg"`
	Advurl       string `json:"advurl"`
	Isrecommand  int    `json:"isrecommand"`
}
