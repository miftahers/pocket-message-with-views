package main

import (
	"log"
	"pocket-message-with-views/config"
	"pocket-message-with-views/route"

	"github.com/gofiber/template/html"
)

func main() {

	config.InitConfig()
	config.InitDB()

	// Load Templates
	engine := html.New("./public/views", ".tmpl")

	// Route app
	app := route.InitRoute(engine)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
