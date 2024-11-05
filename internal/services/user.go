package services

import (
	"errors"
	"gin-api/internal/models"
	"gin-api/internal/repositories"
	"gin-api/internal/requests"
)

type UserService struct {
	repo *repositories.UserRepository
}

// NewUserService 构造函数
func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// GetUsers 获取用户列表服务
func (us *UserService) GetUsers(queryParams *requests.UserQueryParams) ([]models.UserListResponse, int64, error) {
	return us.repo.GetUsers(queryParams)
}

// GetUserInfo 获取单个用户
func (us *UserService) GetUserInfo(userID uint) (*models.User, error) {
	return us.repo.GetUserByField("ID", userID)
}

// CreateUser 创建用户服务
func (us *UserService) CreateUser(user *models.User) error {
	existUser, err := us.repo.GetUserByField("Name", user.Name)
	if err != nil {
		return err
	}
	if existUser != nil {
		return errors.New("用户已存在")
	}
	return us.repo.CreateUser(user)
}

// UpdateUser 更新用户
func (us *UserService) UpdateUser(user *models.User) error {
	// 验证用户是否存在
	existingUser, err := us.repo.GetUserByField("id", user.ID.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("用户不存在")
	}

	// 更新用户信息
	existingUser.Name = user.Name
	existingUser.Password = user.Password
	existingUser.Mobile = user.Mobile
	return us.repo.UpdateUser(existingUser)
}

// DeleteUser 删除用户服务
func (us *UserService) DeleteUser(userID uint) error {
	// 验证用户是否存在
	existingUser, err := us.repo.GetUserByField("id", userID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("用户不存在")
	}
	return us.repo.DeleteUser(userID)
}
