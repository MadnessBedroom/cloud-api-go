package validators

import (
	"cloud-api-go/pkg/database"
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"strings"
)

func init() {
	/*
		自定义规则 not_exists，验证请求数据必须不存在数据库中。
		常用于保证数据库某个字段的值是唯一的，如：用户名、邮箱、手机号...
		not_exists 参数可以用两种，一种是 2 个参数，一种是 3 个参数
		使用方法一："not_exists:username,email"，两个参数
		使用方法二："not_exists:username,email,32"，排除用户 ID 为 32 的用户
	*/
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		// 第一个参数，数据库表的名称，如 gamer
		tableName := rng[0]
		// 第二个参数，字段的名称，如 username
		dbFiled := rng[1]
		// 第三个参数，需要排除的用户 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}
		// 请求的数据
		requestValue := value.(string)
		// 拼接查询 SQL 语句
		query := database.DB.Table(tableName).Where(dbFiled+"=?", requestValue)
		// 如果传了第三个参数，过滤掉第三个参数在数据库对应的那条数据
		if len(exceptID) > 0 {
			query.Where("id!=?", exceptID)
		}
		// 查询数据库
		var count int64
		query.Count(&count)
		// count 不等于 0 时，表示能找到数据
		if count != 0 {
			// 如果有自定义错误消息就使用自定义的
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误信息
			return fmt.Errorf("%v 已被占用", requestValue)
		}

		// 验证通过
		return nil
	})

	// 自定义规则 exists，确保数据库存在某条数据
	// 一个使用场景是创建话题时需要附带 category_id 分类 ID 为参数，此时需要保证
	// category_id 的值在数据库中存在，即可使用：
	// exists:categories,id
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")
		tableName := rng[0]
		dbFiled := rng[1]
		requestValue := value.(string)
		var count int64
		database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue).Count(&count)
		// 验证不通过，数据不存在
		if count == 0 {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 不存在", requestValue)
		}
		return nil
	})
}
