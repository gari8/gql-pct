package model

import "github.com/gari8/gqlgen-pct/domain"

type Place struct {
	ID        string  `gorm:"id"`
	Name      string  `gorm:"name"`
	Latitude  float64 `gorm:"latitude"`
	Longitude float64 `gorm:"longitude"`
	PlaceType int     `gorm:"placeType"`
}

func (p *Place) ToDomain() *domain.Place {
	return &domain.Place{
		ID:        p.ID,
		Name:      p.Name,
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		PlaceType: domain.AllPlaceType[p.PlaceType],
	}
}
