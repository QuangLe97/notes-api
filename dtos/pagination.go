package dtos

type Pagination struct {
	Limit     int         `json:"limit"`
	Page      int         `json:"page"`
	Sort      string      `json:"sort"`
	TotalRows int64       `json:"total_rows"`
	Rows      interface{} `json:"rows"`
	Searchs   []Search    `json:"searchs"`
}
