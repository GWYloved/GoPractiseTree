package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"fmt"
	"os"
	"io"
	"syscall"
)

func main()  {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.POST("/save", save)
	e.POST("/saveAvator", saveAvator)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
func hello(c echo.Context) error  {
	return c.String(http.StatusOK, "Hello, World!")
}

//for /users/:id
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	fmt.Println(id)
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team +", member: "+member+"\n")
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func saveAvator(c echo.Context) error {
	name := c.QueryParam("name")
	avator, err := c.FormFile("avator")
	if err != nil {
		return err
	}
	fmt.Printf(avator.Filename)
	src, err := avator.Open()
	checkNotNil(err)
	defer src.Close()

	dst, err := os.Create(avator.Filename)
	checkNotNil(err)
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thank you ! "+name +"</b>")
}

func checkNotNil(err error) error {
	if err!= nil {
		 return err
	}
	syscall.Exit(101)
	return nil
}
