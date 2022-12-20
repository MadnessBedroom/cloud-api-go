package app

import (
	"cloud-api-go/pkg/config"
	"time"
)

func IsLocal() bool {
	return config.Get("app.env") == "local"
}

func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(chinaTimezone)
}
