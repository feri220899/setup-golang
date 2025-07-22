package testmodule

type TestModule struct {
	Id               uint   `json:"id"`
	Import_file_path string `json:"import_file_path"`
	Import_status    string `json:"import_status"`
	Import_start     int    `json:"import_start"`
	Import_batch     int    `json:"import_batch"`
	Import_total_row int    `json:"import_total_row"`
}

type DataExcel struct {
	Id          uint   `json:"id"`
	Nama_Kolom1 string `json:"nama_kolom1"`
	Nama_Kolom2 string `json:"nama_kolom2"`
	Nama_Kolom3 string `json:"nama_kolom3"`
	Nama_Kolom4 string `json:"nama_kolom4"`
	Nama_Kolom5 string `json:"nama_kolom5"`
	Nama_Kolom6 string `json:"nama_kolom6"`
	Nama_Kolom7 string `json:"nama_kolom7"`
	Nama_Kolom8 string `json:"nama_kolom8"`
	Nama_Kolom9 string `json:"nama_kolom9"`
}
