package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Endpoint para saludar
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "¡Hola, mundo! Esta es mi API (que no hace nada) para probar Travis CI")
	})

	// Iniciar el servidor
	e.Logger.Fatal(e.Start(":8080"))
}
