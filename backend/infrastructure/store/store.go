package store

import (
	"fmt"
	"myapp/config"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/samber/do"
)

func NewStore(i *do.Injector) (*sqlx.DB, error) {
	dsn := mysql.Config{
		DBName:    config.DBName,
		User:      config.DBUser,
		Passwd:    config.DBPassword,
		Addr:      fmt.Sprintf("%s:%d", config.DBHostName, config.DBPort),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.Local,
	}

	return sqlx.Open("mysql", dsn.FormatDSN())
}
