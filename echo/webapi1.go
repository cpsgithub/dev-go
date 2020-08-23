package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type WebUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.POST("/users", saveUser)
	e.POST("/maps", saveMaps)
	e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		fmt.Printf("\n c.Bind is error.\n")
		return err
	}
	fmt.Printf("\n reunt to json.\n")
	return c.JSON(http.StatusOK, m)
}

func saveMaps(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		fmt.Printf("\n c.Bind is error.\n")
		return err
	}
	fmt.Printf("\n reunt to json.\n")
	return c.JSON(http.StatusOK, m)
}

func saveUser(c echo.Context) error {
	u := new(WebUser)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
