package model

import (
	"challenge/api/service"
	"encoding/json"

	"github.com/rafael180496/libcore/database"
	utl "github.com/rafael180496/libcore/utility"
)

type (
	/*User : estructura para inicio de sesion*/
	User struct {
		Username string `json:"user"`
		Password string `json:"password"`
	}
	/*TokenUser : estructura para enviar tokens al api*/
	TokenUser struct {
		User  string `json:"user"`
		Token string `json:"token"`
	}
)

/*Valid : valida los argumento para iniciar sesion*/
func (p *User) Valid() string {
	return utl.ReturnIf(!utl.IsNilArrayStr(p.Password, p.Username), "PET03", "").(string)
}

/*Login : iniciando sesion*/
func (p *User) Login() (string, []byte) {
	cnx := SendDB()
	result, err := cnx.Query(database.StQuery{
		Querie: SendSQL("sql01"),
		Args: map[string]interface{}{
			"USER": p.Username,
			"PASS": p.Password,
		},
	}, 1, true)
	if err != nil {
		service.PrintD("%s", err.Error())
		return "PET04", nil
	}
	cont, erraux := result[0].ToInt("cont")
	if erraux != nil {
		service.PrintD("%s", erraux.Error())
		return "PET04", nil
	}
	if cont <= 0 {
		return "PET03", nil
	}
	token := utl.GenToken(p.Username)
	errorsql := cnx.Exec([]database.StQuery{
		{
			Querie: SendSQL("del01"),
			Args: map[string]interface{}{
				"USER": p.Username,
			},
		}, {
			Querie: SendSQL("ins01"),
			Args: map[string]interface{}{
				"USER":  p.Username,
				"TOKEN": token,
			},
		},
	}, false)
	if errorsql != nil {
		service.PrintD("%s", errorsql.Error())
		return "PET04", nil
	}
	data, errJSON := json.Marshal(&TokenUser{
		Token: token,
		User:  p.Username,
	})
	if errJSON != nil {
		service.PrintD("%s", errJSON.Error())
		return "PET04", nil
	}
	return "", data
}

/*ValidKey :  valida un token unico*/
func ValidKey(key string) bool {
	cnx := SendDB()
	result, err := cnx.Query(database.StQuery{
		Querie: SendSQL("sql03"),
		Args: map[string]interface{}{
			"TOKEN": key,
		},
	}, 1, false)
	if err != nil {
		service.PrintE("%s", err.Error())
		return false
	}
	cont, erraux := result[0].ToInt("cont")
	if erraux != nil {
		service.PrintE("%s", erraux.Error())
		return false
	}
	if cont <= 0 {
		return false
	}
	return true
}

/*Logout : cierra la sesion del usuario*/
func Logout(key string) string {
	cnx := SendDB()
	errorsql := cnx.Exec([]database.StQuery{
		{
			Querie: SendSQL("del02"),
			Args: map[string]interface{}{
				"TOKEN": key,
			},
		},
	}, false)
	if errorsql != nil {
		service.PrintD("%s", errorsql.Error())
		return "PET07"
	}
	return ""
}
