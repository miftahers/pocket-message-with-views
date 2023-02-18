package route

import (
	"pocket-message-with-views/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func InitRoute(engine *html.Engine) *fiber.App {
	// Creating fiber app
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	//Configure app
	app.Static("/", "./public")

	// Routes
	app.Get("/", controller.Index)

	page := app.Group("/page")
	page.Get("/register", controller.PageRegister)
	page.Get("/login", controller.PageLogin)

	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)       //localhost:3000/api/auth/login
	auth.Post("/register", controller.Register) //localhost:3000/api/auth/Register
	return app
}
