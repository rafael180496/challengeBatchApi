package model

import "github.com/rafael180496/libcore/database"

type (
	/*Client : estructura para inserta los csv*/
	Client struct {
		name       string `csv:"name"`
		segment1   bool   `csv:"segment1"`
		segment2   bool   `csv:"segment2"`
		segment3   bool   `csv:"segment3"`
		segment4   bool   `csv:"segment4"`
		platformID int    `csv:"platformId"`
		clientID   int    `csv:"clientId"`
	}
)

/*ConvertMap : convierte un client en map interface*/
func (p *Client) ConvertMap() database.StData {
	return database.StData{
		"name":       p.name,
		"segment1":   p.segment1,
		"segment2":   p.segment2,
		"segment3":   p.segment3,
		"segment4":   p.segment4,
		"platformId": p.platformID,
		"clientId":   p.clientID,
	}
}
