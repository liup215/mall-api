package model

type UniAccount struct {
	Uniacid      int    `gorm:"primary_key;AUTO_INCREMENT" json:"uniacid"`
	Groupid      int    `json:"groupid"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DefaultAcid  int    `json:"default_acid"`
	Rank         int    `json:"rank"`
	TitleInitial string `json:"title_initial"`
}

type Account struct {
	Acid      int    `gorm:"primary_key;AUTO_INCREMENT" json:"acid"`
	Uniacid   int    `json:"uniacid"`
	Hash      string `json:"hash"`
	Type      int    `json:"type"`
	Isconnect int    `json:"isconnect"`
	Isdeleted int    `json:"isdeleted"`
	Endtime   int    `json:"endtime"`
}
