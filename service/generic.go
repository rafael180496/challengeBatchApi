package service

import (
	"github.com/labstack/echo/v4"
)

/*PrintD : Inserta un registro en el log debug*/
func PrintD(format string, args ...interface{}) {
	if Config.Debug {
		LGenD.Printf(format, args...)
	}
}

/*PrintE : Inserta un registro en el log error*/
func PrintE(format string, args ...interface{}) {
	LGenE.Printf(format, args...)
}

/*SendMsjPet : envia un mensaje de la peticiones pre-cargada */
func SendMsjPet(cod string, data interface{}, e echo.Context) error {
	msjs := GetMsjPet()
	return msjs.Send(cod, data, e)
}

/*Valid : valida la estructa y la configura para el servicio*/
func (p *ConfigServer) Valid() error {

	if p.Puerto <= 0 {
		return Msjcore.GetError("AC01")
	}

	return nil
}
