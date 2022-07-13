package env

import "go-gin-template/tool/common_helper"

var (
	DatabaseUrl = common_helper.GetOsEnv("DATABASE_URL", "xxx")
	TokenSecret = common_helper.GetOsEnv("TOKEN_SECRET", "xxx")
)
