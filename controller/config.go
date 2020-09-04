package controller

import (
	"challenge/api/model"
	s "challenge/api/service"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	utl "github.com/rafael180496/libcore/utility"
)

/*StarServer : inicia el servicio.*/
func StarServer(e *echo.Echo) {
	e.Logger.Fatal(e.Start(":" + utl.IntToStr(s.Config.Puerto)))
}

/*serverHeader : Muestra la configuracion del header del proyecto.*/
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headServ := s.APPNAME
		c.Response().Header().Set(echo.HeaderServer, headServ)
		return next(c)
	}
}

/*ConfigServer : configuracion de los servidores*/
func ConfigServer(e *echo.Echo) error {
	e.Use(serverHeader)

	err := configLogPet(e)
	if err != nil {
		return err
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowHeaders, echo.HeaderAuthorization},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:" + echo.HeaderAuthorization,
		Skipper:   skipServer,
		Validator: validServerKey,
	}))

	return nil
}

/*configLogPet : Configura los log de peticiones.*/
func configLogPet(e *echo.Echo) error {
	archlog, err := utl.FileNew(fmt.Sprintf("%s/%s%s.log", s.LOGPET, s.APPNAME, utl.TimetoStr(time.Now())))
	if err != nil {
		return err
	}
	wrt := io.MultiWriter(os.Stdout, archlog)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:           wrt,
		Format:           s.FormatLogger,
		CustomTimeFormat: s.FormatFechaLogger,
	}))
	return nil
}

/*skipServer : skipper de los server  son los path que no van hacer tomados por el auth.*/
func skipServer(c echo.Context) bool {
	return utl.InStr(c.Path(), URLSkipper...)
}

/*ValidServerKey : Valida los token del api rest.*/
func validServerKey(key string, c echo.Context) (bool, error) {
	resp := model.ValidKey(key)
	if !resp {
		return false, s.SendMsjPet("PET02", nil, c)
	}
	return true, nil
}
