package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDB() (db *gorm.DB, err error) {
	// dsn := "host=localhost user=myuser dbname=golang_tutorial port=5432"
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  `user=postgres password=a12345 dbname=DBDEV port=5432 sslmode=disable`,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "BLOG.", // schema name
			SingularTable: false,
		}})
	return
}
