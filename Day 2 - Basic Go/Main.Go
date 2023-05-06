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
	e.GET("/users/:id", UserByIDController)
	e.POST("/login", LoginControler)

	e.Start(":8000")
	}

	func HelloController(c echo.Context) error {
		return c.JSON(200, "Hello, World!")
	}

	func UserController(c echo.Context) error {
		countryParams:= c.QueryParam("country")
		user :=User{
			ID: 1,
			Nama: countryParams,
		}
		return c.JSON(200, user)
	}

	func UserByIDController(c echo.Context) error {
		id := c.Param("id")
		user :=User{
			ID: 1,
			Nama: id,
		}
		return c.JSON(200, user)
	}

	func LoginControler(c echo.Context) error {
		email:= c.FormValue("email")
		password:= c.FormValue("password")
		user :=User{
			ID: 1,
			Nama: email + password,
		}
		return c.JSON(200, user)
	}