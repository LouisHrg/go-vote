package provider

import (
	"github.com/jinzhu/gorm"
	// I don't know this is the way the documentation is
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/go-vote/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {

	conn, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db = conn

	db.LogMode(true)

	db.AutoMigrate(&model.User{}, &model.Survey{}, &model.Response{})

	var newSurvey = &model.Survey{Question: "Plutot java ou .net ?"}

	db.Create(&newSurvey)
	db.Create(&model.Response{Survey: *newSurvey, Message: "Eclatax javax"})
	db.Create(&model.Response{Survey: *newSurvey, Message: "Mdrlolptdr .net"})

	db.Create(&model.User{
		Email:     "vundaboy@gmail.com",
		Firstname: "Louis",
		Lastname:  "Harang",
		Age:       24,
		Password:  "super-secret"})

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
