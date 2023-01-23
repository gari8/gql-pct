package loader

import (
	"context"
	"fmt"
	"github.com/gari8/gqlgen-pct/domain"
	"github.com/graph-gophers/dataloader"
)

func newPlaceLoaderFunc(
	placeRepo PlaceRepo,
) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var ids []string
		for _, key := range keys {
			ids = append(ids, key.String())
		}

		records, err := placeRepo.FindByIDs(ids)
		if err != nil {
			return []*dataloader.Result{}
		}

		placeByID := map[string]*domain.Place{}
		for _, record := range records {
			placeByID[record.ID] = record
		}

		results := make([]*dataloader.Result, len(keys))
		for i, key := range keys {
			k := key.String()
			results[i] = &dataloader.Result{Data: nil, Error: nil}
			if place, ok := placeByID[k]; ok {
				results[i].Data = place
			} else {
				results[i].Error = fmt.Errorf("place[key=%s] not found", k)
			}
		}

		return results
	}
}

func LoadPlace(ctx context.Context, id string) (*domain.Place, error) {
	loader := ctx.Value(loadersKey).(*Loaders)
	thunk := loader.PlaceByID.Load(ctx, dataloader.StringKey(id))
	data, err := thunk()
	if err != nil {
		return nil, err
	}
	return data.(*domain.Place), nil
}
