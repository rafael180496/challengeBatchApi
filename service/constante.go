package service

import (
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rafael180496/libcore/database"
	"github.com/rafael180496/libcore/utility"
)

type (
	/*ConfigServer : configuraciones del servicio*/
	ConfigServer struct {
		Debug  bool `ini:"debug"`
		Puerto int  `ini:"puerto"`
	}
)

var (
	/*Config : configuraciones del server*/
	Config = ConfigServer{
		Debug:  false,
		Puerto: 8001,
	}
	/*LocalCache : intancia para el control de variables cache de la libreria
	https://github.com/patrickmn/go-cache
	5 minutos default
	10 por item
	*/
	LocalCache *cache.Cache = cache.New(5*time.Minute, 10*time.Minute)
	/*KEYMSG : clave de mensaje para guardar en cache*/
	KEYMSG = "MPET"
	/*DbCx : variable de conexion de base de datos general*/
	DbCx database.StConect
	/*LGenE : Variables de log general error*/
	LGenE utility.StLog
	/*LGenD : Variables de log general debug*/
	LGenD utility.StLog
	/*PATHLOG : path donde se guardan los logs de error y debug */
	PATHLOG = CONFIGPATH + "/logs"
	/*LOGPET : path donde se guadan los log de peticiones */
	LOGPET = PATHLOG + "/server"
	/*LOGDEBUG : path donde se guadan los log de debug */
	LOGDEBUG = PATHLOG + "/debug"
	/*LOGERROR : path donde se guadan los log de errores */
	LOGERROR = PATHLOG + "/error"
	/*CONFIGINI : archivo de configuracion del servicio*/
	CONFIGINI = CONFIGPATH + "/" + APPNAME + ".ini"
)

const (
	/*CONFIGPATH : carpeta de configuraciones del servicio*/
	CONFIGPATH = "./config"
	/*APPNAME : nombre del servicio.*/
	APPNAME = "challenge"
	/*server config */

	/*FormatLogger  : Formato configurado para el logger de peticion */
	FormatLogger = `Method[${method}] | Url[${path}] | Status[${status}] | Host[${host}] | Remote_ip[${remote_ip}] | User_agent  : [${user_agent}] | Time[${time_custom}]
`  /*`
	================================================================================================================
	ID          : [${id}]
	Time        : [${time_custom}]
	Status      : [${status}]
	Method      : [${method}]
	Url         : [${path}]
	Host        : [${host}]
	Remote_ip   : [${remote_ip}]
	User_agent  : [${user_agent}]
	`*/
	/*FormatFechaLogger : Formato de fecha de los logs.*/
	FormatFechaLogger = "2006-01-02 15:04:05.00000"
	/*FormatFechaPostgresql : fomati de fecga de la base  de datos de postgresql*/
	FormatFechaPostgresql = "2006-01-02"
)
