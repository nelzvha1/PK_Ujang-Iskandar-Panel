package main

import "github.com/labstack/echo/v4"

type User struct {
	ID int
	Nama string
}

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.GET("/users", UserController)

	e.Start(":8000")
	}

	func HelloController(c echo.Context) error {
		return c.JSON(200, "Hello, World!")
	}

	func UserController(c echo.Context) error {
		user :=User{
			ID: 1,
			Nama: "Nelzvha",
		}
		return c.JSON(200, user)
	}
