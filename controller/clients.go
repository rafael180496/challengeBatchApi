package controller

import (
	"challenge/api/model"
	"challenge/api/service"
	"encoding/json"

	echo "github.com/labstack/echo/v4"
	server "github.com/rafael180496/libcore/server"
)

var (
	/*ClientsController : controlador para clientes*/
	ClientsController = server.Controller{
		Pets: []server.HTTPPet{
			{
				Path: ROUTER["client"],
				Tip:  server.GET,
				Pet: func(c echo.Context) error {
					var (
						clientParam model.ParamClient
						err         error
						v           interface{}
					)
					if err = c.Bind(&clientParam); err != nil {
						return service.SendMsjPet("PET09", nil, c)
					}
					code := clientParam.Valid()
					if code != "" {
						return service.SendMsjPet(code, nil, c)
					}
					result, codeSQL := clientParam.GetClients()
					if codeSQL != "" {
						return service.SendMsjPet(codeSQL, nil, c)
					}
					err = json.Unmarshal(result, &v)
					if err != nil {
						return service.SendMsjPet("PET02", nil, c)
					}
					return service.SendMsjPet("PET00", v, c)
				},
			},
		},
	}
)
