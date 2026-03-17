package handlers

import (
	"projectfiber/models"
	"projectfiber/responses"
	"projectfiber/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPlaylists(c *fiber.Ctx) error {

	playlists, err := services.GetAllPlaylists()
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > len(playlists) {
		limit = len(playlists)
	}

	return responses.Success(c, playlists[:limit])
}

func GetPlaylistByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	playlist, err := services.GetPlaylistByID(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}
	return responses.Success(c, playlist)
}

func CreatePlaylist(c *fiber.Ctx) error {
	var playlist models.Playlists

	if err := c.BodyParser(&playlist); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	createdPlaylist, err := services.CreatePlaylist(playlist)
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	return responses.Success(c, createdPlaylist)
}

func UpdatePlaylist(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	playlist, err := services.UpdatePlaylist(id, data)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, playlist)
}

func FullUpdatePlaylist(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Playlists

	if err := c.BodyParser(&updated); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	updatedPlaylist, err := services.FullUpdatePlaylists(id, updated)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, updatedPlaylist)
}
func DeletePlaylist(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	err = services.DeletePlaylist(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return c.SendStatus(204)
}
