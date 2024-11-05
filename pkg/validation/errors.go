package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

// TranslateError 将校验错误翻译为中文
func TranslateError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var errs []string
		for _, e := range validationErrors {
			errs = append(errs, translate(e))
		}
		return errors.New(strings.Join(errs, "; "))
	}
	return err
}

func translate(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "字段 " + e.Field() + " 是必填项"
	case "min":
		return "字段 " + e.Field() + " 的最小长度为 " + e.Param()
	case "max":
		return "字段 " + e.Field() + " 的最大长度为 " + e.Param()
	case "len":
		return "字段 " + e.Field() + " 的长度必须为 " + e.Param()
	case "email":
		return "字段 " + e.Field() + " 必须是有效的电子邮件地址"
	case "url":
		return "字段 " + e.Field() + " 必须是有效的URL地址"
	case "uuid":
		return "字段 " + e.Field() + " 必须是有效的UUID"
	case "numeric":
		return "字段 " + e.Field() + " 必须是数值"
	case "alpha":
		return "字段 " + e.Field() + " 必须是字母"
	case "alphanum":
		return "字段 " + e.Field() + " 必须是字母或数字"
	case "alphaunicode":
		return "字段 " + e.Field() + " 必须是Unicode字母"
	case "alphanumunicode":
		return "字段 " + e.Field() + " 必须是Unicode字母或数字"
	case "gt":
		return "字段 " + e.Field() + " 必须大于 " + e.Param()
	case "gte":
		return "字段 " + e.Field() + " 必须大于或等于 " + e.Param()
	case "lt":
		return "字段 " + e.Field() + " 必须小于 " + e.Param()
	case "lte":
		return "字段 " + e.Field() + " 必须小于或等于 " + e.Param()
	case "eq":
		return "字段 " + e.Field() + " 必须等于 " + e.Param()
	case "ne":
		return "字段 " + e.Field() + " 必须不等于 " + e.Param()
	case "oneof":
		return "字段 " + e.Field() + " 必须是以下值之一: " + e.Param()
	case "contains":
		return "字段 " + e.Field() + " 必须包含子字符串 " + e.Param()
	case "excludes":
		return "字段 " + e.Field() + " 不能包含子字符串 " + e.Param()
	case "startswith":
		return "字段 " + e.Field() + " 必须以 " + e.Param() + " 开头"
	case "endswith":
		return "字段 " + e.Field() + " 必须以 " + e.Param() + " 结尾"
	case "ip":
		return "字段 " + e.Field() + " 必须是有效的IP地址"
	// 添加更多校验标签的翻译
	default:
		return "字段 " + e.Field() + " 校验失败"
	}
}
