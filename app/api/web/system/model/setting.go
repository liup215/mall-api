package model

type CoreSettings struct {
	Key   string `orm:"key" json:"key"`
	Value string `orm:"value" json:"value"`
}
