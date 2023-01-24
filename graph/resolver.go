package graph

import (
	"github.com/gari8/gqlgen-pct/domain"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type PlaceRepo interface {
	FindAll(placeType *domain.PlaceType) ([]*domain.Place, error)
	FindByID(id string) (*domain.Place, error)
	FindByIDs(ids []string) ([]*domain.Place, error)
}
type ProgramRepo interface {
	FindAll(programType *domain.ProgramType, placeIds []*string) ([]*domain.Program, error)
	FindByID(id string) (*domain.Program, error)
}

type Resolver struct {
	PlaceRepo
	ProgramRepo
}
