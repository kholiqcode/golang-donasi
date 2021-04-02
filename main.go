package main

import (
	"donasi/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/donasi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.Name = "Abdul Kholiq"
	userInput.Email = "kholiqbisnsis@gmail.com"
	userInput.Occupation = "anak band"
	userInput.Password = "asd123asd"

	userService.RegisterUser(userInput)
}
