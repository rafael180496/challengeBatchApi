package model

import (
	"challenge/api/service"
	"encoding/json"

	"github.com/rafael180496/libcore/database"
	"github.com/rafael180496/libcore/utility"
)

type (
	/*FilesUpload : extrae un archivo parametrizado*/
	FilesUpload struct {
		Name string `json:"namefile" form:"namefile" query:"namefile"`
	}
)

/*UploadFile : funcion para subir archivos*/
func UploadFile(name, path string) (string, []byte) {
	cnx := SendDB()
	errorsql := cnx.Exec([]database.StQuery{
		{
			Querie: SendSQL("ins02"),
			Args: map[string]interface{}{
				"FILENAME": name,
				"PATHFILE": path,
			},
		},
	}, false)
	if errorsql != nil {
		service.PrintD("%s", errorsql.Error())
		return "PET13", nil
	}
	data, errJSON := json.Marshal(&FilesUpload{
		Name: name,
	})
	if errJSON != nil {
		service.PrintD("%s", errJSON.Error())
		return "PET13", nil
	}
	return "", data
}

/*DownloadFile : descarga de un archivo*/
func (p *FilesUpload) DownloadFile() (string, string) {
	cnx := SendDB()
	result, err := cnx.Query(database.StQuery{
		Querie: SendSQL("sql04"),
		Args: map[string]interface{}{
			"FILENAME": p.Name,
		},
	}, 1, false)
	if err != nil {
		service.PrintD("%s", err.Error())
		return "PET16", ""
	}
	if len(result) <= 0 {
		return "PET15", ""
	}
	path, _ := result[0].ToString("pathfile")
	if !utility.FileExist(path, false) {
		namefile, _ := result[0].ToString("namefile")
		errorsql := cnx.Exec([]database.StQuery{
			{
				Querie: SendSQL("del03"),
				Args: map[string]interface{}{
					"FILENAME": namefile,
				},
			},
		}, false)
		if errorsql != nil {
			service.PrintD("%s", errorsql.Error())
			return "PET14", ""
		}
		return "PET14", ""
	}
	return "", path
}
