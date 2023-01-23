package model

import "github.com/gari8/gqlgen-pct/domain"

type Program struct {
	ID          string `gorm:"id"`
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	Image       string `gorm:"image"`
	ProgramType int    `gorm:"program_type"`
	PlaceID     string `gorm:"place_id"`
}

func (p *Program) ToDomain() *domain.Program {
	return &domain.Program{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		ProgramType: domain.AllProgramType[p.ProgramType],
		PlaceID:     p.PlaceID,
	}
}
