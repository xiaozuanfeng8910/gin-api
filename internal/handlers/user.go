package handlers

import (
	"fmt"
	"gin-api/internal/models"
	"gin-api/internal/repositories"
	"gin-api/internal/requests"
	"gin-api/internal/services"
	"gin-api/pkg/response"
	"gin-api/pkg/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

type UserHandler struct {
	logger    *zap.Logger
	validator validation.Validator
	db        *gorm.DB
	service   *services.UserService
}

// NewUserHandler 是 UserHandler 的构造函数
func NewUserHandler(logger *zap.Logger, db *gorm.DB, validator validation.Validator) *UserHandler {
	// 创建 UserRepository 和 UserService 实例
	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	return &UserHandler{
		logger:    logger,
		db:        db,
		validator: validator,
		service:   userSvc,
	}
}

// GetUsers 获取用户列表接口
func (uh *UserHandler) GetUsers(c *gin.Context) {
	uh.logger.Info("获取用户列表接口")

	// 绑定查询参数
	var queryParams *requests.UserQueryParams
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		response.Fail(c, 20000, fmt.Sprintf("参数绑定失败:%v", err))
		return
	}

	// 设置默认值
	queryParams.SetDefaults()

	users, total, err := uh.service.GetUsers(queryParams)
	if err != nil {
		response.Fail(c, 20000, fmt.Sprintf("获取用户列表接口失败:%v", err))
		return
	}
	response.Success(c, gin.H{
		"total": total,
		"page":  queryParams.PageSize,
		"users": users,
	}, "成功")

}

// GetUserInfo 获取用户详情
func (uh *UserHandler) GetUserInfo(c *gin.Context) {
	msg := "获取用户详情成功"
	// 获取用户id
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg = fmt.Sprintf("用户ID解析失败:%v", err)
		uh.logger.Error(msg)
		response.ValidateFail(c, msg)
		return
	}

	user, err := uh.service.GetUserInfo(uint(userID))
	if err != nil {
		msg = fmt.Sprintf("获取用户详情失败:%v", err)
		uh.logger.Error(msg)
		customError := response.CustomError{20400, msg}
		response.FailByError(c, customError)
		return
	}

	response.Success(c, user, msg)
	return
}

// CreateUser 创建用户
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user requests.UserCreateRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		uh.logger.Error("数据绑定失败", zap.Error(err))
		response.ValidateFail(c, fmt.Sprintf("参数校验失败:%v", err))
		return
	}

	if err := uh.validator.Validate(user); err != nil {
		uh.logger.Error("数据校验失败", zap.Error(err))
		response.BusinessFail(c, fmt.Sprintf("用户校验失败:%v", err))
		return
	}

	userModel := &models.User{
		Name:     user.Name,
		Password: user.Password,
		Mobile:   user.Mobile,
	}

	if err := uh.service.CreateUser(userModel); err != nil {
		uh.logger.Error("创建用户失败", zap.Error(err))
		customError := response.CustomError{20100, fmt.Sprintf("创建用户失败:%v", err)}
		response.FailByError(c, customError)
		return
	}
	response.Success(c, nil, "成功创建用户")
	return
}

// UpdateUser 更新用户
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	var user requests.UserUpdateRequest
	msg := "用户更新成功"
	// 获取用户id
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg = fmt.Sprintf("用户ID解析失败:%v", err)
		uh.logger.Error(msg)
		response.ValidateFail(c, msg)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		msg = fmt.Sprintf("数据绑定失败:%v", err)
		uh.logger.Error(msg)
		response.ValidateFail(c, msg)
		return
	}
	if err := uh.validator.Validate(user); err != nil {
		msg = fmt.Sprintf("数据校验失败:%v", err)
		uh.logger.Error(msg)
		response.BusinessFail(c, msg)
		return
	}
	userModel := &models.User{
		ID: models.ID{
			ID: uint(userId),
		},
		Name:     user.Name,
		Mobile:   user.Mobile,
		Password: user.Password,
	}
	if err := uh.service.UpdateUser(userModel); err != nil {
		msg = fmt.Sprintf("更新用户失败:%v", err)
		uh.logger.Error(msg)
		customError := response.CustomError{20200, msg}
		response.FailByError(c, customError)
		return
	}
	response.Success(c, nil, msg)
	return
}

// DeleteUser 删除用户
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	msg := "成功删除用户"
	// 获取用户ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg = fmt.Sprintf("用户ID解析失败:%v", err)
		uh.logger.Error(msg)
		response.ValidateFail(c, msg)
		return
	}
	if err := uh.service.DeleteUser(uint(userID)); err != nil {
		msg = fmt.Sprintf("删除用户失败:%v", err)
		uh.logger.Error(msg)
		customError := response.CustomError{20300, msg}
		response.FailByError(c, customError)
		return
	}
	response.Success(c, nil, msg)
	return
}
