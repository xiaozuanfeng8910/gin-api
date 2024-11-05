package utils

import "os"

// GetEnv 获取环境变量值
func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// PathExists 判断目录是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
