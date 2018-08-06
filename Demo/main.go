package main

import (
	"log"

	"github.com/NlaakStudios/gowaf/app"
	"github.com/NlaakStudios/gowaf/controller"
	"github.com/NlaakStudios/gowaf/models"
	// Import Application Models and Controllers here
)

// Register holds all the Models and Controllers you wish to register for the WebApp
func Register(a *app.App) {
	// Register models
	a.Model.Register(
		//Framework Common Models
		&models.Example{},
		&models.Address{},
		&models.Email{},
		&models.Gender{},
		&models.Note{},
		&models.PersonName{},
		&models.PersonType{},
		&models.Phone{},
		&models.Company{},
		&models.Person{},
		&models.Account{},
	)
	a.Model.LogMode(true)
	a.Model.AutoMigrateAll()

	// Register Controllers
	a.AddController(controller.NewExample)
	a.AddController(controller.NewLanding)
	a.AddController(controller.NewAccount)
	a.AddController(controller.NewAddress)
	a.AddController(controller.NewEmail)
	a.AddController(controller.NewGender)
	a.AddController(controller.NewPersonName)
	a.AddController(controller.NewPersonType)
	a.AddController(controller.NewPhone)
	a.AddController(controller.NewNote)
	a.AddController(controller.NewCompany)
	a.AddController(controller.NewPerson)
}

func main() {
	// Create a new MVC Application object
	app, err := app.NewMVC(Version())
	if err != nil {
		log.Fatal(err)
	} else {
		app.Run(Register)
	}
}
