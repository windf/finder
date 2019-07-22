package model

type Record struct {
	ID          int64  `json:"id" gorm:"column:id"`
	PublisherId int64  `json:"publisher_id" gorm:"column:publisher_id"`
	Name        string `json:"name" gorm:"column:name"`
	Sex         int    `json:"sex" gorm:"column:sex"`
	Photo       string `json:"photo" gorm:"column:photo"`
	Address     string `json:"address" gorm:"column:address"`
	Date        string `json:"date" gorm:"column:date"`
	Remark      string `json:"remark" gorm:"column:remark"`
	Status      int    `json:"status" gorm:"column:status"`
	CreateTime  int64  `json:"create_time" gorm:"column:create_time"`
	UpdateTime  int64  `json:"update_time" gorm:"column:update_time"`
}
