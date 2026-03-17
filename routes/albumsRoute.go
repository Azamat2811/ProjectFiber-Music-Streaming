package routes

import (
	"projectfiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func AlbumRoutes(app *fiber.App) {
	albums := app.Group("/albums")

	albums.Post("/", handlers.CreateAlbum)
	albums.Get("/", handlers.GetAlbums)
	albums.Get("/:id", handlers.GetAlbumByID)
	albums.Patch("/:id", handlers.UpdateAlbum)
	albums.Put("/:id", handlers.FullUpdateAlbums)
	albums.Delete("/:id", handlers.DeleteAlbum)
	albums.Get("/:id/tracks", handlers.GetTracksByAlbumID)

}
