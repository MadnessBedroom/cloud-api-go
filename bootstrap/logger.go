package bootstrap

import (
	"cloud-api-go/pkg/config"
	"cloud-api-go/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.type"),
		config.GetString("log.level"),
	)
}
