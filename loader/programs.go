package loader

import (
	"context"
	"github.com/gari8/gqlgen-pct/domain"
	"github.com/graph-gophers/dataloader"
)

func newProgramsLoaderFunc(
	programRepo ProgramRepo,
) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var ids []*string
		for _, key := range keys {
			k := key.String()
			ids = append(ids, &k)
		}

		records, err := programRepo.FindAll(nil, ids)
		if err != nil {
			return []*dataloader.Result{}
		}

		programByID := map[string][]*domain.Program{}
		for _, record := range records {
			programByID[record.ID] = append(programByID[record.ID], record)
		}

		programs := make([][]*domain.Program, len(ids))
		for i, id := range ids {
			programs[i] = programByID[*id]
		}

		results := make([]*dataloader.Result, len(programs))
		for i := range programs {
			results[i] = &dataloader.Result{Data: programs[i], Error: nil}
		}

		return results
	}
}

func LoadPrograms(ctx context.Context, id string, programType *domain.ProgramType) ([]*domain.Program, error) {
	loader := ctx.Value(loadersKey).(*Loaders)
	thunk := loader.ProgramsByID.Load(ctx, dataloader.StringKey(id))
	data, err := thunk()
	if err != nil {
		return nil, err
	}
	var programs []*domain.Program
	for _, p := range data.([]*domain.Program) {
		if programType != nil {
			if p.ProgramType != *programType {
				continue
			}
		}
		programs = append(programs, p)
	}
	return programs, nil
}
