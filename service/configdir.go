package service

import (
	"github.com/go-ini/ini"
	utl "github.com/rafael180496/libcore/utility"
)

/*ListArch : lista de archivo que debe de validar o crear */
var ListArch = utl.StArchMa{
	Archs: []utl.StArch{
		{
			Path:   CONFIGPATH,
			IndDir: true,
		},
		{
			Path:   LOGDEBUG,
			IndDir: true,
		},
		{
			Path:   LOGERROR,
			IndDir: true,
		},
		{
			Path:   LOGPET,
			IndDir: true,
		},
		{
			Path:   CONFIGINI,
			IndDir: false,
		},
	},
}

/*IniDir :inicia y valida si estan las carpeta si no las crea*/
func IniDir() error {
	return ListArch.Create()
}

/*ReadIni : leer el archivo de configuracion */
func ReadIni() error {
	err := DbCx.ConfigINI(CONFIGINI)
	if err != nil {
		return nil
	}
	cfg, erraux := ini.Load(CONFIGINI)
	if erraux != nil {
		return erraux
	}
	err = cfg.Section("server").MapTo(&Config)
	if err != nil {
		return err
	}
	return Config.Valid()
}
