package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()
	app.Get("/:userid/detail", GetByUserId)
	app.Get("/follower/:username", GetByUsername)
	app.Listen(":"+ os.Getenv("PORT"))
	// app.Listen(":8989")

}

type DataUser struct {
	UserId    string `json:"userid"`
	Username  string `json:"username"`
	Followers int    `json:"followers"`
}

var dataUser = []DataUser{
	{UserId: "sammy", Username: "SammyShark", Followers: 987},
	{UserId: "jesse", Username: "JesseOctopus", Followers: 432},
	{UserId: "jamie", Username: "JamieMantisShrimp", Followers: 654},
	{UserId: "drew", Username: "DrewSquid", Followers: 321},

}

func GetByUserId(c *fiber.Ctx) error {
	userId := c.Params("userid") 
	for _, data := range dataUser {
		if data.UserId == userId {
			return c.Status(fiber.StatusCreated).JSON((fiber.Map{
				"Status": 200,
				"data": fiber.Map{
					"username": data.Username,
					"follower": data.Followers,
				},
			}))
		}
	}

	// if data not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "data not found",
	})

}

func GetByUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	for _, data := range dataUser {
		if data.Username == username {
			return c.Status(fiber.StatusCreated).JSON((fiber.Map{
				"Status": 200,
				"data": fiber.Map{
					"follower": data.Followers,
				},
			}))
		}
	}

	// if data not available
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "data not found",
	})

}
