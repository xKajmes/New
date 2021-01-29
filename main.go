package main

import (
	"fmt"
	"new/config"
	"new/mappings"
	"new/models"

	//"github.com/twinj/uuid"
	//"github.com/go-redis/redis/v7"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// C:\Go\src %USERPROFILE%
	// setx GOPATH=$HOME/go
)

// func init() {
// 	//Initializing redis
// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 	   dsn = "localhost:6379"
// 	}
// 	auth.Client = redis.NewClient(&redis.Options{
// 	   Addr: dsn, //redis port
// 	})
// 	_, err := auth.Client.Ping().Result()
// 	if err != nil {
// 	   panic(err)
// 	}
//   }
func main() {

	r := mappings.HTML()	// Mappowanie zmiennych
	db, err := gorm.Open(sqlite.Open("./baza.db"), &gorm.Config{})
	if err != nil {
	  panic("Brak połaczenia z bazą :( ")
	}
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Class{})
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Grade{})
	db.AutoMigrate(&models.Verify{})
	dziekan := &models.User{}
	dziekan.Login = "Dziekan"
	dziekan.Password, _ = config.HashPassword("12345")
	dziekan.Role = "dziekan"
	dziekan.Verify = true
	szukaj := db.Where("login = ?", dziekan.Login).Find(&dziekan)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
	db.Create(&dziekan)
	}
	admin := &models.User{}
	admin.Login = "Admin"
	admin.Password, _ = config.HashPassword("12345")
	admin.Role = "admin"
	admin.Verify = true
	szukaj = db.Where("login = ?", admin.Login).Find(&admin)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
	db.Create(&admin)
	}
	nauczyciel := &models.User{}
	nauczyciel.Login = "Nauczyciel"
	nauczyciel.Password, _ = config.HashPassword("12345")
	nauczyciel.Role = "teacher"
	nauczyciel.Verify = true
	szukaj = db.Where("login = ?", nauczyciel.Login).Find(&nauczyciel)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
	db.Create(&nauczyciel)
	}

	class := &models.Class{}
	class.Name = "1TOI"
	class.ID_Tutor = 3
	szukaj = db.Where("name = ?", class.Name).Find(&class)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
	db.Create(&class)
	}
	user := &models.User{}
	user.Login = "Kamil"
	user.Password, _ = config.HashPassword("12345")
	user.Verify = true
	user2 := &models.Student{}
	user2.ID = 4
	user2.Name = "Kamil"
	user2.Surname = "Ziemann"
	user2.Age = "18"
	user2.ID_Class = 1
	szukaj = db.Where("login = ?", user.Login).Find(&user)	// Jeżeli użytkownik istnieje
	if(szukaj.RowsAffected==0){
	db.Create(&user)
	db.Create(&user2)
	}

	grade := &models.Grade{}
	grade.ID_Teacher = 3
	grade.ID_Student = 4
	grade.Number = 4
	grade.Comment = "ZA PRACE DOMOWĄ KURWA"
	db.Create(&grade)


	config.DB = db
	r.Run(":8080") 	// Odpalenie na porcie :8080
	fmt.Printf("Hello World")
}



//var  client *redis.Client


//   Auth.client = redis.NewClient(&redis.Options{
//      Addr: dsn, //redis port
//   })
//   _, err := Auth.client.Ping().Result()
//   if err != nil {
//      panic(err)
//   }




