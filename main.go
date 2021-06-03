package main

import (
	"database/sql"
	"os"

	. "github.com/amrali/golang-blog/pkg/db"
	"gorm.io/gorm"
	r "github.com/amrali/golang-blog/pkg/router"
)

func main() {

	var sqlDB *sql.DB
	var db *gorm.DB
	{
		var err error
		db, err = ConnectDB()
		if err != nil {
			os.Exit(-1)
		}
		sqlDB, _ = db.DB()
	}
	defer sqlDB.Close()
	//Migrate(db)
	r.HttpRequests()

}
