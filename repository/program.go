package repository

import (
	"github.com/gari8/gqlgen-pct/domain"
	"github.com/gari8/gqlgen-pct/repository/model"
	"gorm.io/gorm"
)

type ProgramRepository struct {
	*gorm.DB
}

func NewProgramRepository(db *gorm.DB) *ProgramRepository {
	return &ProgramRepository{db}
}

func (r ProgramRepository) FindAll(programType *domain.ProgramType, placeIds []*string) ([]*domain.Program, error) {
	var programs []*model.Program
	db := r.DB
	if programType != nil {
		db = db.Where("program_type = ?", programType.ToInt())
	}
	var err error
	if placeIds != nil {
		err = db.Find(&programs, "place_id IN ?", placeIds).Error
	} else {
		err = db.Find(&programs).Error
	}
	if err != nil {
		return nil, err
	}
	var result []*domain.Program
	for _, program := range programs {
		result = append(result, program.ToDomain())
	}
	return result, nil
}

func (r ProgramRepository) FindByID(id string) (*domain.Program, error) {
	var program model.Program
	if err := r.DB.Where("id = ?", id).First(&program).Error; err != nil {
		return nil, err
	}
	return program.ToDomain(), nil
}
