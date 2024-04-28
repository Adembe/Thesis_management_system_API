package database

import (
	"fmt"
	"log"
	"time"

	"go-rest-api/database/migrations"
	"go-rest-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	//str := "host=localhost port=5432 user=postgres dbname=thesis_db sslmode=disable password=123"
	str := "postgres://thesis_db_twk5_user:A5LHBU1nVTWlwsSGR2Lk5XG9RCU0Hcyp@dpg-con3rssf7o1s73fd1ka0-a.singapore-postgres.render.com/thesis_db_twk5"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database


	var p models.User

	p.Email = "admin@gmail.com"
	p.Password = "123"
	p.Fname = "admin"
	p.Type = 1
	p.Programm = 0
	p.Lname = "admin"

	err = db.Create(&p).Error
	if err != nil {
		fmt.Print(err)
	}


	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
