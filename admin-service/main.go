package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("hello")
	server := echo.New()

	// Start Server
	server.Logger.Fatal(server.Start(":8000"))
}
