package entities

import "github.com/jinzhu/gorm"

type Dungeon struct {
	gorm.Model
	TileNode
	Level    int
	Progress int
}
