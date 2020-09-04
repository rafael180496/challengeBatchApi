package process

import (
	"time"

	cr "challenge/api/controller"
	"challenge/api/model"
	m "challenge/api/model"
	s "challenge/api/service"

	"github.com/labstack/echo/v4"
	sv "github.com/rafael180496/libcore/server"
	utl "github.com/rafael180496/libcore/utility"
)

/*Mainprocess : proceso principal del servicio */
func Mainprocess() error {
	api := echo.New()
	model.SetQuerie()
	/*Validando Carpetas de configuracion.*/
	utl.PrintPc(utl.Green, s.Msjcore.GetString("GE02"))
	err := s.IniDir()
	if err != nil {
		utl.PrintRed("ERROR:%s\n", err.Error())
		return err
	}
	utl.PrintPc(utl.Green, s.Msjcore.GetString("GE08"))
	err = s.ReadIni()
	if err != nil {
		return err
	}
	utl.PrintPc(utl.Green, s.Msjcore.GetString("GE09"))
	cnx := m.SendDB()
	if !cnx.Test() {
		err = cnx.ExecBackup()
		if err != nil {
			return err
		}
	}
	utl.PrintPc(utl.Green, s.Msjcore.GetString("GE12"))
	err = s.GetMensaje()
	if err != nil {
		return err
	}
	err = iniciarLogs()
	utl.PrintPc(utl.Green, s.Msjcore.GetString("GE10"))
	if err != nil {
		return err
	}

	err = iniciarServicio(api)
	if err != nil {
		return err
	}
	return nil
}

func iniciarServicio(e *echo.Echo) error {
	err := cr.ConfigServer(e)
	if err != nil {
		return err
	}
	err = sv.AsigServer(e, cr.RouterIndex)
	if err != nil {
		return err
	}
	cr.StarServer(e)
	return nil
}

func iniciarLogs() error {
	if s.Config.Debug {
		s.LGenD.Dir = s.LOGDEBUG
		s.LGenD.Fe = time.Now()
		s.LGenD.Name = s.APPNAME
		s.LGenD.Prefix = "DEBUG"
		err := s.LGenD.Init()
		if err != nil {
			return err
		}
	}
	s.LGenE.Dir = s.LOGERROR
	s.LGenE.Fe = time.Now()
	s.LGenE.Name = s.APPNAME
	s.LGenE.Prefix = "ERROR"
	err := s.LGenE.Init()
	if err != nil {
		return err
	}
	return nil
}
