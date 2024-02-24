package main

import "johnnyjacob/battlesnake/version"

func GetSnakeMeta() *SnakeMeta {

	return &SnakeMeta{
		ApiVersion: 0,
		Author:     "johnnyjacob",
		Color:      "#FF00FF",
		Head:       "default",
		Tail:       "default",
		Version:    version.Version,
	}
}
