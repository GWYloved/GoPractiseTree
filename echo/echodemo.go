package main

import {
"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
}
func main()  {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
func hello(c echo.Context) error  {
	return c.
}
