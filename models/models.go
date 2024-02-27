package models

type SnakeMeta struct {
	ApiVersion int    `json:"apiVersion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
	Version    string `json:"version"`
}

type Coords struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Board struct {
	Height  int      `json:"height"`
	Width   int      `json:"width"`
	Food    []Coords `json:"food"`
	Hazards []Coords `json:"hazards"`
	Snakes  []Snake  `json:"snakes"`
}
type Customizations struct {
	Color string
	Head  string
	Tail  string
}

type Snake struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	Health         int            `json:"health"`
	Body           []Coords       `json:"body"`
	Latency        string         `json:"latency"`
	Head           Coords         `json:"head"`
	Length         int            `json:"length"`
	Shout          string         `json:"shout"`
	Squad          int            `json:"squad"`
	Customizations Customizations `json:"customizations"`
}

type RulesetSettings struct {
}

type RuleSet struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Settings []RulesetSettings `json:"settings"`
}

type Game struct {
	Id      string  `json:"id"`
	Ruleset RuleSet `json:"ruleset"`
	MapType string  `json:"map"`
	TimeOut int     `json:"timeout"`
	Source  string  `json:"source"`
}

type MoveRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type Direction string

const (
	MOVE_UP    Direction = "up"
	MOVE_DOWN  Direction = "down"
	MOVE_LEFT  Direction = "left"
	MOVE_RIGHT Direction = "right"
)

type MoveResponse struct {
	Move  Direction `json:"move"`
	Shout string    `json:"string"`
}
