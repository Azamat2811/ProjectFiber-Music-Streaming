package handlers

import (
	"projectfiber/models"
	"projectfiber/responses"
	"projectfiber/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateAlbum(c *fiber.Ctx) error {
	var album models.Albums

	if err := c.BodyParser(&album); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	createdAlbum, err := services.CreateAlbum(album)
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	return c.Status(201).JSON(createdAlbum)
}

func GetAlbumByID(c *fiber.Ctx) error {
	idStr := c.Params("id")        //idStr = "3"
	id, err := strconv.Atoi(idStr) //id = 3, err != nil
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	album, err := services.GetAlbumByID(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}
	return responses.Success(c, album)
}

func GetAlbums(c *fiber.Ctx) error {

	albums, err := services.GetAllAlbums()
	if err != nil {
		return responses.Error(c, 500, err.Error())
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > len(albums) {
		limit = len(albums)
	}

	return responses.Success(c, albums[:limit])
}

func UpdateAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}

	album, err := services.UpdateAlbum(id, data)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, album)
}

func FullUpdateAlbums(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updated models.Albums

	if err := c.BodyParser(&updated); err != nil {
		return responses.Error(c, 400, "invalid JSON")
	}
	updatedAlbum, err := services.FullUpdateAlbums(id, updated)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, updatedAlbum)
}
func DeleteAlbum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return responses.Error(c, 400, "invalid id")
	}
	err = services.DeleteAlbum(id)
	if err != nil {
		return responses.Error(c, 404, err.Error())
	}

	return responses.Success(c, "album deleted successfully")
}

