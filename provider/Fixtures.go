package provider

import (
	"time"
	"github.com/go-vote/model"
)

func loadFixtures() {
	input := "2018-10-01"
	layout := "2006-01-02"
	start, _ := time.Parse(layout, input)

	input = "2020-10-01"
	layout = "2006-01-02"
	end, _ := time.Parse(layout, input)

	var newSurvey = &model.Survey{
		Title: "Propreté des trottoirs",
		Desc: "Dans le budget qui sera soumis au vote des conseillers de Paris lundi, 32 M€ seront consacrés à l’achat de nouvelles machines, plus silencieuses, plus propres et plus performantes. Il n’y aura pas d’embauche d’agents supplémentaires.",
		StartDate: start,
		EndDate: end,
	}

	db.Create(&newSurvey)
	db.Create(&model.Response{Survey: *newSurvey, Message: "Oui je suis d'accord"})
	db.Create(&model.Response{Survey: *newSurvey, Message: "Non je ne suis pas d'accord"})

	input = "1996-02-08"
	layout = "2006-01-02"
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
		Accesslevel:       2,
		Dateofbirth:	t,
		Password:  "test"})
}
