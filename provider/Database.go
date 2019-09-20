package provider

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	// I don't know this is the way the documentation is
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/go-vote/model"
	"time"
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

	db.AutoMigrate(&model.User{}, &model.Survey{}, &model.Response{})

	var newSurvey = &model.Survey{Question: "Plutot java ou .net ?"}

	db.Create(&newSurvey)
	db.Create(&model.Response{Survey: *newSurvey, Message: "Eclatax javax"})
	db.Create(&model.Response{Survey: *newSurvey, Message: "Eclatax .net"})

	input := "1996-02-08"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	
	db.Create(&model.User{
		Email:     "admin@admin.com",
		Firstname: "Jean",
		Lastname:  "Admin",
		Accesslevel:       1,
		Dateofbirth:	t,
		Password:  "admin"})

	db.Create(&model.User{
		Email:     "test@test.com",
		Firstname: "Test",
		Lastname:  "Test",
		Accesslevel:       1,
		Dateofbirth:	t,
		Password:  "secret"})
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
