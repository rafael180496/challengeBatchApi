package controller

import (
	"challenge/api/model"
	"challenge/api/service"
	s "challenge/api/service"
	"encoding/json"
	"fmt"
	"io"
	"os"

	echo "github.com/labstack/echo/v4"
	server "github.com/rafael180496/libcore/server"
	"github.com/rafael180496/libcore/utility"
)

var (
	/*FilesController :  control de archivos*/
	FilesController = server.Controller{
		Pets: []server.HTTPPet{
			{
				Path: ROUTER["upload"],
				Tip:  server.POST,
				Pet: func(c echo.Context) error {
					var v interface{}
					file, err := c.FormFile("file")
					if err != nil {
						service.PrintE("%s", err.Error())
						return s.SendMsjPet("PET11", nil, c)
					}
					src, err := file.Open()
					if err != nil {
						service.PrintE("%s", err.Error())
						return s.SendMsjPet("PET12", nil, c)
					}
					dst, err := os.Create(fmt.Sprintf("%s/%s", service.SRCFILES, file.Filename))
					if err != nil {
						service.PrintE("%s", err.Error())
						return s.SendMsjPet("PET13", nil, c)
					}
					if _, err = io.Copy(dst, src); err != nil {
						service.PrintE("%s", err.Error())
						return s.SendMsjPet("PET14", nil, c)
					}

					code, resp := model.UploadFile(file.Filename, dst.Name())
					src.Close()
					dst.Close()
					if code != "" {
						utility.RmFile(dst.Name())
						return s.SendMsjPet(code, nil, c)
					}
					err = json.Unmarshal(resp, &v)
					if err != nil {
						return service.SendMsjPet("PET06", nil, c)
					}
					return service.SendMsjPet("PET00", v, c)
				},
			}, {
				Path: ROUTER["dow"],
				Tip:  server.GET,
				Pet: func(c echo.Context) error {
					var (
						fileParams model.FilesUpload
						err        error
					)
					if err = c.Bind(&fileParams); err != nil {
						return service.SendMsjPet("PET05", nil, c)
					}
					code, path := fileParams.DownloadFile()
					if code != "" {
						return s.SendMsjPet(code, nil, c)
					}
					return c.File(path)
				},
			},
		},
	}
)
