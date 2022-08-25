package WebsiteModel

type User struct {
	Username string `json:"username" gorm:"column:username;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (User) TableName() string {
	return "users"
}

type Filter struct {
	Username string `json:"username" form:"username"`
}

func (Filter) TableName() string {
	return "users"
}
