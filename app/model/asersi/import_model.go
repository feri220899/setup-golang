package asersimodel

type ImportStatus struct {
	Id               uint   `json:"id"`
	Import_file_path string `json:"import_file_path"`
	Import_status    string `json:"import_status"`
	Import_start     int    `json:"import_start"`
	Import_batch     int    `json:"import_batch"`
	Import_total_row int    `json:"import_total_row"`
	User_id          uint   `json:"User_id"`
}

func (ImportStatus) TableName() string {
	return "import_status"
}
