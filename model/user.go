package model

const (
	UserError = 0
	USER      = 1 //用户
	REVIEWER  = 2 //志愿者
	ADMIN     = 3 //管理员

	StatusOK  = 0
	StatusERR = 1
)

type User struct {
	ID         int64  `json:"id" gorm:"column:id"`
	UserName   string `json:"username" gorm:"column:username"`
	Password   string `json:"password" gorm:"column:password"`
	NickName   string `json:"nickname" gorm:"column:nickname"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Email      string `json:"email" gorm:"column:email"`
	Remark     string `json:"remark" gorm:"column:remark"`
	Status     int    `json:"status" gorm:"column:status"`
	Role       int    `json:"role" gorm:"column:role"`
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}
