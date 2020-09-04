package model

import "github.com/rafael180496/libcore/database"

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
)

/*ConvertMap : convierte un client en map interface*/
func (p *Client) ConvertMap() database.StData {
	return database.StData{
		"name":       p.Name,
		"segment1":   p.Segment1,
		"segment2":   p.Segment2,
		"segment3":   p.Segment3,
		"segment4":   p.Segment4,
		"platformId": p.PlatformID,
		"clientId":   p.ClientID,
	}
}
