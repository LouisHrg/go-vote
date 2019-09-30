package provider

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/go-vote/model"
	"github.com/jinzhu/gorm"
	// I don't know this is the way the documentation is
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/validations"
)

var db *gorm.DB

func init() {

	conn, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db = conn

	db.LogMode(true)

	validations.RegisterCallbacks(db)

	db.AutoMigrate(
		&model.User{},
	 	&model.Survey{},
	 	&model.Response{},
	 	&model.IP{},
	)

	//loadFixtures()
}

// GetDB : getter of the instance of the database
func GetDB() *gorm.DB {
	return db
}

// GetAll : gets all the class data in a paginated way
func GetAll(objects interface{}, page, limit int) interface{} {

	paginator := pagination.Paging(&pagination.Param{
		DB:      db.Set("gorm:auto_preload", true),
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"created_at asc"},
		ShowSQL: false,
	}, objects)

	return paginator
}
