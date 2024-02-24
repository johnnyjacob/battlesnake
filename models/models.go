package models

// {
// 	"apiversion": "1",
// 	"author": "MyUsername",
// 	"color": "#888888",
// 	"head": "default",
// 	"tail": "default",
// 	"version": "0.0.1-beta"
//   }

type SnakeMeta struct {
	ApiVersion int    `json:"apiVersion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}
