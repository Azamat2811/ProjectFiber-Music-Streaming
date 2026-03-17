package main

import (
	"log"
	"projectfiber/db"
	"projectfiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	app.Use(logMiddleware)

	routes.UserRoutes(app)
	routes.AlbumRoutes(app)
	routes.TrackRoutes(app)
	routes.PlaylistRoutes(app)
	routes.ArtistRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func logMiddleware(c *fiber.Ctx) error {
	log.Printf("[%s] %s request:%s", c.Method(), c.Path(), c.Body())
	return c.Next()
}
