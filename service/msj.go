package service

import (
	"github.com/patrickmn/go-cache"
	ser "github.com/rafael180496/libcore/server"
	"github.com/rafael180496/libcore/utility"
)

type (

	/*MaMjsJSON : captura los json de los mensajes*/
	MaMjsJSON struct {
		Msjs []MjsJSON `json:"msjs"`
	}
	/*MjsJSON : captura los json de los mensajes*/
	MjsJSON struct {
		Cod  string `json:"cod"`
		Code int    `json:"code"`
		Msj  string `json:"msj"`
	}
)

/*GetMensaje : captura lo mensajes y lo guarda en la base de datos*/
func GetMensaje() error {
	datamsj := MaMjsJSON{
		[]MjsJSON{
			{
				Cod:  "PET00",
				Code: 200,
				Msj:  "Correcto",
			},
			{
				Cod:  "PET01",
				Code: 200,
				Msj:  "Api Challenge activo",
			}, {
				Cod:  "PET02",
				Code: 401,
				Msj:  "El servicio no esta autorizado.",
			}, {
				Cod:  "PET03",
				Code: 412,
				Msj:  "Clave o Usuario incorrecto.",
			},
			{
				Cod:  "PET04",
				Code: 500,
				Msj:  "Error al obtener usuario",
			},
			{
				Cod:  "PET05",
				Code: 412,
				Msj:  "Error al convertir json de entrada",
			},
			{
				Cod:  "PET06",
				Code: 500,
				Msj:  "Error al convertir json.",
			}, {
				Cod:  "PET07",
				Code: 500,
				Msj:  "Error en cerrar la sesion",
			},
		},
	}
	SetMsjPet(datamsj.toMapMsjPet())
	return nil
}

/*GetMsjPet : Obtiene los mensajes guardado en cache */
func GetMsjPet() ser.MapMsjPet {
	if data, ok := LocalCache.Get(KEYMSG); ok {
		base := data.(ser.MapMsjPet)
		return base
	}
	return ser.MapMsjPet{}
}

/*SetMsjPet : guarda los mensajes del proyecto en cache*/
func SetMsjPet(datos ser.MapMsjPet) {
	LocalCache.Set(KEYMSG, datos, cache.NoExpiration)
}

/*toMapMsjPet : tramforma los mensajes a MapMsjPet*/
func (p *MaMjsJSON) toMapMsjPet() ser.MapMsjPet {
	msjs := make(ser.MapMsjPet)
	for _, item := range p.Msjs {
		cod, m := item.toStMsjPet()
		msjs[cod] = m
	}
	return msjs
}

/*toStMsjPet : tramforma los mensajes a StMsjPet*/
func (p *MjsJSON) toStMsjPet() (string, ser.StMsjPet) {
	return p.Cod, ser.StMsjPet{Code: p.Code, Msj: p.Msj}
}

var (
	/*Msjcore : mensajes de error en las librerias*/
	Msjcore = utility.StMsj{
		Store: map[string]string{
			"GE01": "Mensaje no encontrado.",
			"GE02": "-> Validando carpetas de configuración\n",
			"GE08": "-> Validando Archivo de configuración\n",
			"GE09": "-> Probando base de datos\n",
			"GE10": "-> Iniciando Servicio\n",
			"GE12": "-> Cargando los mensajes\n",
			"AC01": "Puerto no valido.",
			"AC02": "Protocolo no válido {http,https}",
			"PC02": "[X] Sin conexión a base de datos\n",
			"BT01": "No existe el comando que ejecuto",
		},
	}
)
