package locality

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	LocalityCode Code
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("locality code '%v' not found", e.LocalityCode)
}

type Repository interface {
	CreateLocality(ctx context.Context, locality *Locality) error
	RetrieveLocality(ctx context.Context, code Code) (*Locality, error)
	UpdateLocality(
		ctx context.Context,
		code Code,
		updateFn func(ctx context.Context, tr *Locality) (*Locality, error),
	) error
	DeleteLocality(ctx context.Context, locality *Locality) error
}
