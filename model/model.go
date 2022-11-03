package model

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Short struct {
	gorm.Model
	ID     uint64 `json:"id" gorm:"primaryKey;auto_increment"`
	Url    string `json:"url" gorm:"not null;default:null"`
	Short  string `json:"short" gorm:"not null;default:null"`
	Clicks uint64 `json:"clicks" gorm:"not null;default:0"`
}

var db *gorm.DB

func Init() {
	var err error

	db, err = gorm.Open(sqlite.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Short{})

	if err != nil {
		fmt.Println(err)
	}
}
