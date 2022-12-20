package config

import "cloud-api-go/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name":     config.Env("APP_NAME", "GoCloud"),
			"env":      config.Env("APP_ENV", "local"),
			"debug":    config.Env("APP_DEBUG", false),
			"port":     config.Env("APP_PORT", "3000"),
			"key":      config.Env("APP_KEY", "80f7032065753e289ff1d7ef21ea3ae5"),
			"url":      config.Env("APP_URL", "http://localhost:3000"),
			"timezone": config.Env("TIMEZONE", "Asis/Shanghai"),
			//"api_domain": config.Env("TIMEZONE", "Asis/Shanghai"),
		}
	})
}
