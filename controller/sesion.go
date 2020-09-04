package controller

import (
	"challenge/api/model"
	"challenge/api/service"
	"encoding/json"

	echo "github.com/labstack/echo/v4"
	server "github.com/rafael180496/libcore/server"
)

var (
	/*SesionController : inicio de sesion */
	SesionController = server.Controller{
		Pets: []server.HTTPPet{
			{
				Path: ROUTER["login"],
				Tip:  server.POST,
				Pet: func(c echo.Context) error {
					var (
						sesionIn model.User
						err      error
						v        interface{}
					)
					if err = c.Bind(&sesionIn); err != nil {
						return service.SendMsjPet("PET09", nil, c)
					}
					code := sesionIn.Valid()
					if code != "" {
						return service.SendMsjPet(code, nil, c)
					}
					codeSQL, result := sesionIn.Login()
					if codeSQL != "" {
						return service.SendMsjPet(codeSQL, nil, c)
					}
					err = json.Unmarshal(result, &v)
					if err != nil {
						return service.SendMsjPet("PET02", nil, c)
					}
					return service.SendMsjPet("PET00", v, c)
				},
			}, {
				Path: ROUTER["logout"],
				Tip:  server.GET,
				Pet: func(c echo.Context) error {
					var (
						v interface{}
					)
					key := server.ExtKey(c)
					code := model.Logout(key)
					if code != "" {
						return service.SendMsjPet(code, nil, c)
					}
					return service.SendMsjPet("PET00", v, c)
				},
			},
		},
	}
)
