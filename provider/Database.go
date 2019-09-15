package provider


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-vote/model"
)

var db *gorm.DB

func init() {

  conn, err := gorm.Open("sqlite3", "test.db")

  if err != nil {
    panic("failed to connect database")
  }

  db = conn
  db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
