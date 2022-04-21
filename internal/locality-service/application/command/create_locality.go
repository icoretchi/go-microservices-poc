package command

import (
	"context"
	"github.com/icoretchi/go-microservices-poc/internal/common/logs"
	"github.com/icoretchi/go-microservices-poc/internal/locality/domain/locality"
)

type CreateLocality struct {
	Code            locality.Code
	StatisticalCode locality.StatisticalCode
	Name            locality.Name
	Status          locality.Status
}

type CreateLocalityHandler struct {
	repo locality.Repository
}

func NewCreateLocalityHandler(localityRepo locality.Repository) CreateLocalityHandler {
	if localityRepo == nil {
		panic("nil localityRepo")
	}

	return CreateLocalityHandler{localityRepo}
}

func (h CreateLocalityHandler) Handle(ctx context.Context, cmd CreateLocality) error {
	defer func() {
		logs.LogCommandExecution("ScheduleTraining", cmd, err)
	}()

	lt, err := locality.NewLocality(cmd.Code, cmd.StatisticalCode, cmd.Name, cmd.Status)
	if err != nil {
		return err
	}

	if err := h.repo.CreateLocality(ctx, lt); err != nil {
		return err
	}

	return nil
}
