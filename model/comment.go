package model

type Comment struct {
	ID         int64  `json:"id" gorm:"column:id"`
	RecordId   int64  `json:"record_id" gorm:"column:record_id"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Photo      string `json:"photo" gorm:"column:photo"`
	Remark     string `json:"remark" gorm:"column:remark"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	ShowTime   string `json:"show_time"`
}
