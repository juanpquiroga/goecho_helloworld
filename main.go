package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type User struct {
    Name  string `json:"name" xml:"name" form:"name" query:"name"`
    Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// Basic
func helloWorld(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
}

// Path Param
func getUser(c echo.Context) error {
    id := c.Param("id")
    return c.String(http.StatusOK, id)
}

// Query Param
func show(c echo.Context) error {
    // Get team and member from the query string
    team := c.QueryParam("team")
    member := c.QueryParam("member")
    return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}

func save(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, u)
}

func main() {
    e := echo.New()
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
    e.GET("/", helloWorld)
    e.GET("/users/:id", getUser)
    e.GET("/show", show)
    e.POST("/users", save )
    e.Logger.Fatal(e.Start(":1323"))
}

