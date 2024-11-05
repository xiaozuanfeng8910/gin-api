package validation

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Validator 校验器接口
type Validator interface {
	Validate(data interface{}) error
}

// CustomValidator 自定义校验器
type CustomValidator struct {
	logger   *zap.Logger
	validate *validator.Validate
}

// NewCustomValidator 是 CustomValidator 的构造函数
func NewCustomValidator(logger *zap.Logger) *CustomValidator {
	validate := validator.New()
	return &CustomValidator{
		logger:   logger,
		validate: validate,
	}
}

// Validate 实现接口Validator
func (cv *CustomValidator) Validate(data interface{}) error {
	err := cv.validate.Struct(data)
	if err != nil {
		cv.logger.Error("数据校验失败", zap.Error(err))
		return TranslateError(err)
	}
	cv.logger.Info("数据校验通过")
	return nil
}
