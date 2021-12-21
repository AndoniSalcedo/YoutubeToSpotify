package createplaylist

type Body struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string        `json:"href"`
	ID     string        `json:"id"`
	Images []interface{} `json:"images"`
	Name   string        `json:"name"`
	Owner  struct {
		DisplayName  string `json:"display_name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"owner"`
	PrimaryColor interface{} `json:"primary_color"`
	Public       bool        `json:"public"`
	SnapshotID   string      `json:"snapshot_id"`
	Tracks       struct {
		Href     string        `json:"href"`
		Items    []interface{} `json:"items"`
		Limit    int           `json:"limit"`
		Next     interface{}   `json:"next"`
		Offset   int           `json:"offset"`
		Previous interface{}   `json:"previous"`
		Total    int           `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}
