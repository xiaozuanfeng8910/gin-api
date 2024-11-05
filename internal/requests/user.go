package requests

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Mobile   string `json:"mobile" validate:"required,numeric,len=11"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Password string `json:"password,omitempty"`
}

// UserQueryParams 查询参数结构体
type UserQueryParams struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Name     string `json:"name" form:"name"`
	Age      int    `json:"age" form:"age"`
}

// SetDefaults 设置默认值
func (q *UserQueryParams) SetDefaults() {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 10
	}
}
