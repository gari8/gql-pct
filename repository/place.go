package repository

import (
	"github.com/gari8/gqlgen-pct/domain"
	"github.com/gari8/gqlgen-pct/repository/model"
	"gorm.io/gorm"
)

type PlaceRepository struct {
	*gorm.DB
}

func NewPlaceRepository(db *gorm.DB) *PlaceRepository {
	return &PlaceRepository{db}
}

func (r PlaceRepository) FindAll(placeType *domain.PlaceType) ([]*domain.Place, error) {
	var places []*model.Place
	db := r.DB
	if placeType != nil {
		db = db.Where("place_type = ?", placeType.ToInt())
	}
	if err := db.Find(&places).Error; err != nil {
		return nil, err
	}
	var result []*domain.Place
	for _, place := range places {
		result = append(result, place.ToDomain())
	}
	return result, nil
}

func (r PlaceRepository) FindByID(id string) (*domain.Place, error) {
	var place model.Place
	if err := r.DB.Where("id = ?", id).First(&place).Error; err != nil {
		return nil, err
	}
	return place.ToDomain(), nil
}

func (r PlaceRepository) FindByIDs(ids []string) ([]*domain.Place, error) {
	var records []*model.Place
	if err := r.DB.Find(&records, ids).Error; err != nil {
		return nil, err
	}
	var places []*domain.Place
	for _, record := range records {
		places = append(places, record.ToDomain())
	}
	return places, nil
}
