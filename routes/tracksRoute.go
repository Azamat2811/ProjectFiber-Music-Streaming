package routes

import (
	"projectfiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func TrackRoutes(app *fiber.App) {

	tracks := app.Group("/tracks")

	tracks.Post("/", handlers.CreateTrack)
	tracks.Get("/", handlers.GetTracks)
	tracks.Get("/:id", handlers.GetTrackByID)
	tracks.Patch("/:id", handlers.UpdateTrack)
	tracks.Put("/:id", handlers.FullUpdateTrack)
	tracks.Delete("/:id", handlers.DeleteTrack)
}
