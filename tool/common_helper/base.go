package common_helper

import "os"

//获取环境变量
func GetOsEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}
