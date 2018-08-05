package main

import (
	"log"

	"github.com/NlaakStudios/gowaf/app"
	"github.com/NlaakStudios/gowaf/controller"
	"github.com/NlaakStudios/gowaf/models"
	// Import Application Models and Controllers here
)

func main() {
	// Create a new MVC Application object
	app, err := app.NewMVC()
	if err != nil {
		log.Fatal(err)
	}

	// Register models
	app.Model.Register(
		//Framework Common Models
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
	app.Model.LogMode(true)
	app.Model.AutoMigrateAll()

	// Register Controllers
	app.AddController(controller.NewExample)
	app.AddController(controller.NewLanding)
	app.AddController(controller.NewAccount)
	app.AddController(controller.NewAddress)
	app.AddController(controller.NewEmail)
	app.AddController(controller.NewGender)
	app.AddController(controller.NewPersonName)
	app.AddController(controller.NewPersonType)
	app.AddController(controller.NewPhone)
	app.AddController(controller.NewNote)
	app.AddController(controller.NewCompany)
	app.AddController(controller.NewPerson)

	//Run the app
	app.Run()
}
