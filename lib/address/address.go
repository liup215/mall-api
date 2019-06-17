package address

type Address struct {
	Province string    `bson:"province,omitempty" json:"province,omitempty"`
	City     string    `bson:"city,omitempty" json:"city,omitempty"`
	District string    `bson:"district,omitempty" json:"district,omitempty"`
	Detail   string    `bson:"detail,omitempty" json:"detail,omitempty"`
	Location []float64 `bson:"location,omitempty" json:"location,omitempty"`
}
