package models

type SnakeMeta struct {
	ApiVersion int    `json:"apiVersion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}
