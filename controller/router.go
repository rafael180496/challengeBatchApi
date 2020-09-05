package controller

import "github.com/rafael180496/libcore/server"

var (
	/*ROUTER :  path de todo el proyecto core*/
	ROUTER = map[string]string{
		/*path de prueba */
		"test":   "/",
		"login":  "/login",
		"logout": "/logout",
		"client": "/clients"}
	/*URLSkipper : skipper pesonalizados del controlador.*/
	URLSkipper = []string{
		ROUTER["test"],
		ROUTER["login"],
	}

	/*RouterIndex : router maestro de para administrar los controladores
	 */
	RouterIndex = []server.Controller{
		PruebaController,
		SesionController,
		ClientsController,
	}
)
