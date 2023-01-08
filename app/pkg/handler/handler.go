package handler

import (
	"fmt"
	"github.com/atadzan/url-shortener/app/model"
	"github.com/atadzan/url-shortener/app/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Redirect(c *fiber.Ctx) error {
	golyUrl := c.Params("redirect")
	goly, err := model.FindByGolyUrl(golyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "can't find goly in db" + err.Error(),
		})
	}
	// grab any stats you want...
	goly.Clicked += 1
	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

func GetAllGolies(c *fiber.Ctx) error {
	golies, err := model.GetAllGolies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

func GetGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error colud not parse id" + err.Error(),
		})
	}
	goly, err := model.GetGoly(id)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly from db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error passing JSON" + err.Error(),
		})
	}
	if goly.Random {
		goly.Goly = utils.RandomURL(8)
	}
	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not create goly in db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func UpdateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var goly model.Goly

	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse json" + err.Error(),
		})
	}

	err = model.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not update goly link in db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

func DeleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not parse id from url" + err.Error(),
		})
	}
	err = model.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not delete from db" + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted",
	})
}
