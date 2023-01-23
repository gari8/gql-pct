package loader

import (
	"context"
	"fmt"
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

		programByID := map[string]*domain.Program{}
		for _, record := range records {
			programByID[record.ID] = record
		}

		results := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			k := key.String()
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			if place, ok := programByID[k]; ok {
				results[i].Data = place
			} else {
				results[i].Error = fmt.Errorf("program[key=%s] not found", k)
			}
		}

		return results
	}
}

func LoadPrograms(ctx context.Context, id string, programType *domain.ProgramType) ([]*domain.Program, error) {
	loader := ctx.Value(loadersKey).(*Loaders)
	thunk := loader.ProgramsByID.LoadMany(ctx, dataloader.Keys{dataloader.StringKey(id)})
	data, errors := thunk()
	if errors != nil {
		return nil, errors[0]
	}
	var programs []*domain.Program
	for _, d := range data {
		p := d.(*domain.Program)
		if programType != nil {
			if p.ProgramType != *programType {
				continue
			}
		}
		programs = append(programs, p)
	}
	return programs, nil
}
