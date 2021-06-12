package user

import (
	"fmt"

	"github.com/GeovaniiSilva/go-fiber-tutorial/viperenv"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	gorm.Model

	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" gorm:"unique"`
}

var host = viperenv.ViperEnvVariable("host")
var admin = viperenv.ViperEnvVariable("admin")
var password = viperenv.ViperEnvVariable("password")

var DNS = fmt.Sprintf("%s:%s@%s", admin, password, host)

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}

	DB.AutoMigrate(&User{})
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)

	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	DB.Find(&user, id)

	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(500).SendString("User not Available")
	}

	DB.Delete(&user)
	return c.SendString("User Deleted!")
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if res := DB.Create(&user); res.Error != nil {
		c.Status(500).SendString("User email already exists!")
	}

	return c.JSON(&user)

}

func UpdateUser(c *fiber.Ctx) error {

	id := c.Params("id")

	user := new(User)
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User does not exist")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return c.JSON(&user)
}
