package dumymodel

type Dumy struct {
	Id    uint   `json:"id"`
	Bulan int    `json:"bulan"`
	Tahun int    `json:"tahun"`
	Dumy  string `json:"dumy"`
}

func (Dumy) TableName() string {
	return "data_dumy"
}
