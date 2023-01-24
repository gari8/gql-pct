package loader

import (
	"context"
	"github.com/gari8/gqlgen-pct/domain"
	"github.com/graph-gophers/dataloader"
	"net/http"
	"time"
)

type loadersKeyType string

const loadersKey loadersKeyType = "loaders"

type Loaders struct {
	PlaceByID    *dataloader.Loader
	ProgramsByID *dataloader.Loader
}

type PlaceRepo interface {
	FindByIDs(ids []string) ([]*domain.Place, error)
}

type ProgramRepo interface {
	FindAll(programType *domain.ProgramType, placeIds []*string) ([]*domain.Program, error)
}

func DataLoaderMiddleware(
	placeRepo PlaceRepo,
	programRepo ProgramRepo,
) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
				PlaceByID:    newBatchedFunc(newPlaceLoaderFunc(placeRepo)),
				ProgramsByID: newBatchedFunc(newProgramsLoaderFunc(programRepo)),
			})
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func NewLoaders(
	placeRepo PlaceRepo,
	programRepo ProgramRepo,
) *Loaders {
	return &Loaders{
		PlaceByID:    newBatchedFunc(newPlaceLoaderFunc(placeRepo)),
		ProgramsByID: newBatchedFunc(newProgramsLoaderFunc(programRepo)),
	}
}

func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func newBatchedFunc(bf dataloader.BatchFunc) *dataloader.Loader {
	return dataloader.NewBatchedLoader(bf, dataloader.WithWait(1*time.Millisecond))
}
