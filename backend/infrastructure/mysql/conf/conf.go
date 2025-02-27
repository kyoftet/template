package mysqlconf

import (
	"backend/infrastructure/mysql/ent"
	"backend/util"
	"fmt"
)

func GetDB() (*ent.Client, error) {
	dsn := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?charset=utf8&parseTime=True",
		util.GetEnv("MYSQL_USER", "user"),
		util.GetEnv("MYSQL_PASSWORD", "password"),
		util.GetEnv("MYSQL_HOST", "localhost"),
		util.GetEnv("MYSQL_PORT", "3306"),
		util.GetEnv("MYSQL_NAME", "template"),
	)
	return ent.Open("mysql", dsn)
}
