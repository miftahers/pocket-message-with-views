package controller

import (
	"pocket-message-with-views/config"
	"pocket-message-with-views/helper"
	"pocket-message-with-views/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(fiber.Map{
				"isError":      true,
				"errorMessage": "username atau password tidak valid",
			})
		} else {
			return c.JSON(fiber.Map{
				"isError":      true,
				"errorMessage": "Kesalahan server, silahkan coba lagi",
			})
		}
	}

	valid := helper.CheckPasswordHash(password, user.Password)
	if !valid {
		return c.JSON(fiber.Map{
			"isError":      true,
			"errorMessage": "username atau password tidak valid",
		})
	}

	return c.Render("dashboard", fiber.Map{
		"data": user,
	})
}

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	user.Username = username

	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	user.Password = hashedPassword
	user.UUID = uuid.New()
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := config.DB.Create(&user).Error; err != nil {
				return fiber.ErrInternalServerError
			}
		} else {
			return fiber.ErrInternalServerError
		}
	}
	return c.Render("loginPage", nil)
}
