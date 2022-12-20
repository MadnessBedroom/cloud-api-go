package config

import (
	"cloud-api-go/pkg/helpers"
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
	"os"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {
	// 初始化 viper 库
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appEnv")
	viper.AutomaticEnv()
	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string) {
	loadEnv(env)
	loadConfig()
}

// loadEnv 加载环境变量文件
func loadEnv(envSuffix string) {
	// 默认加载项目根目录下的 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filePath := ".env" + envPath
		if _, err := os.Stat(filePath); err == nil {
			envPath = filePath
		}
	}
	// 设置配置文件的文件名称
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监控配置文件，变更时重新加载
	viper.WatchConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn)
	}
}

func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

func internalGet(path string, defaultValue ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}

	return viper.Get(path)
}

// Add 新增配置
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
