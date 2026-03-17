package routes

import (
	"projectfiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func ArtistRoutes(app *fiber.App) {

	artists := app.Group("/artists")

	artists.Post("/", handlers.CreateArtist)
	artists.Get("/", handlers.GetArtists)
	artists.Get("/:id", handlers.GetArtistByID)
	artists.Patch("/:id", handlers.UpdateArtist)
	artists.Put("/:id", handlers.FullUpdateArtist)
	artists.Delete("/:id", handlers.DeleteArtist)
	artists.Get("/:id/albums", handlers.GetAlbumsByArtist)

}
