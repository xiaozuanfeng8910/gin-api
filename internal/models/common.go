package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Model 的定义
type Model struct {
	ID
	Timestamps
	SoftDeletes
}
