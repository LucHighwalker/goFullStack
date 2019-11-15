package magic

import (
	"golang-starter-pack/model"
)

type MagicStore interface {
	GetByName(string) (*model.Card, error)
	CreateCard(*model.Card) error
	UpdateCard(*model.Card, []string) error
	DeleteCard(*model.Card) error
	List(offset, limit int) ([]model.Card, int, error)
}
