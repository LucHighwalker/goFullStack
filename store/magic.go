package store

import (
	"github.com/jinzhu/gorm"
	"golang-starter-pack/model"
)

type MagicStore struct {
	db *gorm.DB
}

func NewMagicStore(db *gorm.DB) *MagicStore {
	return &MagicStore{
		db: db,
	}
}

func (ms *MagicStore) GetByName(n string) (*model.Card, error) {
	var m model.Card
	err := ms.db.Where(&model.Card{Name: n}).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (as *MagicStore) CreateCard(m *model.Card) error {
	tx := as.db.Begin()
	if err := tx.Create(&m).Error; err != nil {
		return err
	}
	if err := tx.Where(m.ID).Find(&m).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *MagicStore) UpdateCard(m *model.Card, tagList []string) error {
	tx := as.db.Begin()
	if err := tx.Model(m).Update(m).Error; err != nil {
		return err
	}
	return tx.Commit().Error
}

func (as *MagicStore) DeleteCard(m *model.Card) error {
	return as.db.Delete(m).Error
}

func (as *MagicStore) List(offset, limit int) ([]model.Card, int, error) {
	var (
		cards []model.Card
		count int
	)
	as.db.Model(&cards).Count(&count)
	as.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&cards)
	return cards, count, nil
}
