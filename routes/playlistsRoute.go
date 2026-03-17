package routes

import (
	"projectfiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func PlaylistRoutes(app *fiber.App) {

	playlist := app.Group("/playlists")

	playlist.Post("/", handlers.CreatePlaylist)
	playlist.Get("/", handlers.GetPlaylists)
	playlist.Get("/:id", handlers.GetPlaylistByID)
	playlist.Patch("/:id", handlers.UpdatePlaylist)
	playlist.Put("/:id", handlers.FullUpdatePlaylist)
	playlist.Delete("/:id", handlers.DeletePlaylist)
}
