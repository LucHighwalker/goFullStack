package model

import (
	"github.com/jinzhu/gorm"
)

type CardType struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type CardEffect struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Effect string `gorm:"not null"`
}

type ManaColor struct {
	gorm.Model
	Color int `gorm:"not null"`
}

type Card struct {
	gorm.Model
	Name   string     `gorm:"unique_index;not null"`
	Type   []CardType `gorm:"not null"`
	Effect []CardEffect
	Colors []ManaColor
	Flavor string
	Image  string `gorm:"not null"`
}
