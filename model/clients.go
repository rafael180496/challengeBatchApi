package model

import (
	"challenge/api/service"
	"fmt"
	"strings"

	db "github.com/rafael180496/libcore/database"
	utl "github.com/rafael180496/libcore/utility"
)

type (
	/*Client : estructura para inserta los csv*/
	Client struct {
		Name       string `csv:"name"`
		Segment1   bool   `csv:"segment1"`
		Segment2   bool   `csv:"segment2"`
		Segment3   bool   `csv:"segment3"`
		Segment4   bool   `csv:"segment4"`
		PlatformID int    `csv:"platformId"`
		ClientID   int    `csv:"clientId"`
	}
	/*ParamClient : parametros para envio de datos del sql dinamico*/
	ParamClient struct {
		Sort      string `json:"sort" form:"sort" query:"sort"`
		SortField string `json:"sortfield" form:"sortfield" query:"sortfield"`
		Limit     int    `json:"limit" form:"limit" query:"limit"`
		Init      int    `json:"init" form:"init" query:"init"`
	}
)

/*Valid : valida los parametros para sql dinamico*/
func (p *ParamClient) Valid() string {
	p.Init = utl.ReturnIf(p.Init <= 0, 0, p.Init).(int)
	p.Limit = utl.ReturnIf(p.Limit <= 10, 10, p.Limit).(int)
	p.Sort = utl.ReturnIf(!utl.IsNilStr(p.Sort), ASC, p.Sort).(string)
	p.SortField = utl.ReturnIf(!utl.IsNilStr(p.SortField), "name", p.SortField).(string)
	if !utl.InStr(strings.ToLower(p.Sort), ASC, DESC) {
		return "PET08"
	}
	if !utl.InStr(strings.ToLower(p.SortField), Colums...) {
		return "PET08"
	}
	return ""
}

/*GetClients : obtiene los clientes del api*/
func (p *ParamClient) GetClients() ([]byte, string) {
	sqltmp := fmt.Sprintf(SendSQL("sql02"), p.SortField, p.Sort)
	cnx := SendDB()
	resp, err := cnx.QueryJSON(
		db.StQuery{
			Querie: sqltmp,
			Args: map[string]interface{}{
				"CANT": p.Limit,
				"INI":  p.Init,
			},
		}, 0, false, false)
	if err != nil {
		service.PrintE("%s", err.Error())
		return nil, "PET10"
	}
	return resp, ""
}

/*ConvertMap : convierte un client en map interface*/
func (p *Client) ConvertMap() db.StData {
	return db.StData{
		"name":       p.Name,
		"segment1":   p.Segment1,
		"segment2":   p.Segment2,
		"segment3":   p.Segment3,
		"segment4":   p.Segment4,
		"platformId": p.PlatformID,
		"clientId":   p.ClientID,
	}
}
