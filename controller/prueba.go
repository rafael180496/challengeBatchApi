package controller

import (
	s "challenge/api/service"

	echo "github.com/labstack/echo/v4"
	server "github.com/rafael180496/libcore/server"
)

var (
	/*PruebaController : prueba de peticion. */
	PruebaController = server.Controller{
		Pets: []server.HTTPPet{
			{
				Path: ROUTER["test"],
				Tip:  server.GET,
				Pet: func(c echo.Context) error {
					return s.SendMsjPet("PET01", nil, c)
				},
			},
		},
	}
)
