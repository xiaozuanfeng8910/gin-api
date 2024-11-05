package repositories

import (
	"gin-api/internal/models"
	"gin-api/internal/requests"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 构造函数
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetUsers 分页获取用户列表
func (ur *UserRepository) GetUsers(queryParams *requests.UserQueryParams) ([]models.UserListResponse, int64, error) {
	var users []models.UserListResponse
	var total int64

	// 构建查询条件
	query := ur.db.Model(&models.User{})
	name := queryParams.Name
	age := queryParams.Age
	page := queryParams.Page
	pageSize := queryParams.PageSize

	if name != "" {
		query = query.Where("name like ?", "%"+name+"%")
	}
	if age > 0 {
		query = query.Where("age = ?", age)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 9, err
	}
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.
		Offset(offset).Limit(pageSize).Find(&users).
		Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// GetUserByField 根据字段名和值获取用户
func (ur *UserRepository) GetUserByField(field string, value interface{}) (*models.User, error) {
	var user models.User
	if err := ur.db.Where(field+" = ?", value).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

// UpdateUser 更新用户
func (ur *UserRepository) UpdateUser(user *models.User) error {
	updateFields := map[string]interface{}{}
	if user.Name != "" {
		updateFields["name"] = user.Name
	}
	if user.Mobile != "" {
		updateFields["mobile"] = user.Mobile
	}
	if user.Password != "" {
		updateFields["password"] = user.Password
	}
	return ur.db.Model(&models.User{}).
		Select("name", "mobile", "password").
		Where("id = ?", user.ID.ID).
		Updates(updateFields).Error
}

// DeleteUser 删除用户
func (ur *UserRepository) DeleteUser(id uint) error {
	return ur.db.Delete(&models.User{}, id).Error
}
