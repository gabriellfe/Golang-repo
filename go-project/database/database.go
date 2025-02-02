package database

import (
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetupDatabase() {
	slog.Info("Conecting to database")
	dsn := "root:root@tcp(127.0.0.1:3306)/usuariodb?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "usuariodb", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,        // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	sqldb, _ := db.DB()
	err := sqldb.Ping()
	if err != nil {
		panic(err)
	}

	slog.Info("Conected to database")
}
