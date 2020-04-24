package gorecipe

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Resolver struct{
     DB *gorm.DB
}


