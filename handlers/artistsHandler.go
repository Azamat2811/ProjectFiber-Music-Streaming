package handlers

import (
	"projectfiber/models"
	"projectfiber/responses"
	"projectfiber/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetArtists(c *fiber.Ctx) error {

	artists, err := services.GetAllArtists()
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > len(artists) {
		limit = len(artists)
	}

	return responses.Success(c, artists[:limit])
}

func GetArtistByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	artist, err := services.GetArtistByID(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, artist)
}

func CreateArtist(c *fiber.Ctx) error {
	var artist models.Artists

	if err := c.BodyParser(&artist); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	createdArtist, err := services.CreateArtist(artist)
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	return responses.Success(c, createdArtist)
}

func UpdateArtist(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	artist, err := services.UpdateArtist(id, data)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, artist)
}

func FullUpdateArtist(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Artists

	if err := c.BodyParser(&updated); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	updatedArtist, err := services.FullUpdateArtists(id, updated)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, updatedArtist)
}
func DeleteArtist(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}
	err = services.DeleteArtist(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return c.SendStatus(204)
}

func GetAlbumsByArtist(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}
	albums, err := services.GetAlbumsByArtist(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return c.JSON(albums)
}
