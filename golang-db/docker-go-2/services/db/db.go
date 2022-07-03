package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB
var err error

func initialMigration() {
	log.Println("Migration initialization")

	db, err = gorm.Open("sqlite3", "/database/test.db")
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
}

func allUsersDB() []User {
	log.Println("allUsersDB handler invoked")

	db, err = gorm.Open("sqlite3", "/database/test.db")
	if err != nil {
		log.Fatalln("allUsersDB --- Couldn't connect to the database:", err)
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	log.Println(fmt.Sprintf("%d users fetched", len(users)))
	return users
}

func newUserDB(name, email string) string {
	log.Println("newUserDB handler invoked")

	db, err = gorm.Open("sqlite3", "/database/test.db")
	if err != nil {
		log.Println("newUserDB --- Couldn't connect to the database:", err)
	}
	defer db.Close()

	newUser := &User{
		Name:  name,
		Email: email,
	}
	db.Create(newUser)

	logMessage := fmt.Sprintf(
		"New User '%s' with email '%s' succesfully created", newUser.Name, newUser.Email)
	log.Println(logMessage)

	return logMessage
}

func deleteUserDB(name string) string {
	log.Println("deleteUserDB handler invoked")

	db, err = gorm.Open("sqlite3", "/database/test.db")
	if err != nil {
		log.Fatalln("deleteUserDB --- Couldn't connect to the database:", err)
	}
	defer db.Close()

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	logMessage := fmt.Sprintf(
		"User '%s' with email '%s' successfully deleted", user.Name, user.Email)
	log.Println(logMessage)

	return logMessage
}

func updateUserDB(name, email string) string {
	log.Println("updateUserDB handler invoked")
	db, err = gorm.Open("sqlite3", "/database/test.db")
	if err != nil {
		log.Fatalln("updateUserDB --- Couldn't connect to the database:", err)
	}
	defer db.Close()

	var user User
	db.Where("name = ?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	logMessage := fmt.Sprintf(
		"User '%s' email updated: %s", user.Name, user.Email)
	log.Println(logMessage)

	return logMessage
}
