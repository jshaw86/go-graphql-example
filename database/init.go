package database

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/vividcortex/mysqlerr"
)

func InitDB(c *Config) *gorm.DB {
	db, err := FetchConnection(c)

	if driverErr, ok := err.(*mysql.MySQLError); ok {
		if driverErr.Number == mysqlerr.ER_BAD_DB_ERROR {
			connstr := fmt.Sprintf("%s:%s@tcp(%s:3306)/?parseTime=True", c.Username, c.Password, c.Hostname)
			db, err = gorm.Open("mysql", connstr)
			// Create the database. This is a one-time step.
			// Comment out if running multiple times - You may see an error otherwise
			db.Exec("CREATE DATABASE test_db")
			db.Exec("USE test_db")
		}
	}

	db.LogMode(true)

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&TodoList{}, &Item{})

	return db
}

func FetchConnection(c *Config) (*gorm.DB, error) {
	connstr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=True", c.Username, c.Password, c.Hostname, c.Database)
	return gorm.Open("mysql", connstr)

}
