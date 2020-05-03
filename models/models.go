package models

import(
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
    DatabaseType string
    Hostname string
    Username string
    Password string
    Database string

}

type List struct {
    ID       int    `json:"id" gorm:"primary_key"`
    Name     string `json:"name"`
    Items []Item `json:"items"`
}

type Item struct {
    ID       int    `json:"id" gorm:"primary_key"`
    ListId int `json:"listId"`
    Name     string `json:"name"`
    DueDate int    `json:"dueDate"`
}

func FetchConnection(c Config) *gorm.DB{
    connstr := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.Username,c.Password,c.Hostname, c.Database)
	db,err := gorm.Open("mysql",connstr)
	if err != nil{
		panic(err)
	}
	return db
}

