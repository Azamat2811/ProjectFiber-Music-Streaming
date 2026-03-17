package handlers

import (
	"projectfiber/models"
	"projectfiber/responses"
	"projectfiber/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetTracks(c *fiber.Ctx) error {

	tracks, err := services.GetAllTracks()
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > len(tracks) {
		limit = len(tracks)
	}

	return responses.Success(c, tracks[:limit])
}

func GetTrackByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	track, err := services.GetTrackByID(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}
	return responses.Success(c, track)
}

func CreateTrack(c *fiber.Ctx) error {
	var track models.Tracks

	if err := c.BodyParser(&track); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	createdTrack, err := services.CreateTrack(track)
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	return responses.Success(c, createdTrack)
}

func UpdateTrack(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	track, err := services.UpdateTrack(id, data)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, track)
}

func FullUpdateTrack(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Tracks

	if err := c.BodyParser(&updated); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	updatedTrack, err := services.FullUpdateTracks(id, updated)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, updatedTrack)
}
func DeleteTrack(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}
	err = services.DeleteTrack(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return c.SendStatus(204)
}

func GetTracksByAlbumID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}
	tracks, err := services.GetTracksByAlbum(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}
	return c.JSON(tracks)
}
