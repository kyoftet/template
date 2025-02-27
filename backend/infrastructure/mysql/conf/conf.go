package mysqlconf

import (
	"backend/infrastructure/mysql/ent"
	"backend/util"
	"github.com/go-sql-driver/mysql"
)

func GetDB() (*ent.Client, error) {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())

	c := mysql.Config{
		User:                 util.GetEnv("MYSQL_USER", ""),
		Passwd:               util.GetEnv("MYSQL_PASSWORD", ""),
		Net:                  "tcp",
		Addr:                 util.GetEnv("MYSQL_HOST", "") + ":" + util.GetEnv("MYSQL_PORT", ""),
		DBName:               util.GetEnv("MYSQL_DATABASE", ""),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	return ent.Open("mysql", c.FormatDSN(), entOptions...)
}
