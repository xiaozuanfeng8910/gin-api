package models

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:手机号"`
	Password string `json:"password" gorm:"not null;default:'';comment:密码"`
	Timestamps
	SoftDeletes
}

type UserListResponse struct {
	ID
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}
