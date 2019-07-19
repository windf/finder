package model

type User struct {
	ID         int64  `json:"id" gorm:"column:id"`
	UserName   int64  `json:"username" gorm:"column:username"`
	Password   string `json:"password" gorm:"column:password"`
	NickName   string `json:"nickname" gorm:"column:nickname"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Email      string `json:"email" gorm:"column:email"`
	Remark     string `json:"remark" gorm:"column:remark"`
	Status     byte   `json:"status" gorm:"column:status"`
	Role       byte   `json:"role" gorm:"column:role"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}
