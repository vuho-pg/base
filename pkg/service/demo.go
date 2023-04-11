package service

import (
	"context"
	"github.com/vuho-pg/base/pkg/dto"
	"github.com/vuho-pg/base/pkg/mapper"
	"github.com/vuho-pg/base/pkg/model"
	"github.com/vuho-pg/base/pkg/repository"
	"github.com/vuho-pg/base/pkg/util/api"
)

type DemoService interface {
}

type demoService struct {
	demoRepo repository.DemoRepository
}

func NewDemoService(demoRepo repository.DemoRepository) DemoService {
	return &demoService{
		demoRepo: demoRepo,
	}
}

func (s *demoService) Create(ctx context.Context, req dto.CreateDemo) (api.Response, error) {
	data := &model.Demo{
		Audit: model.Audit{},
		Name:  req.Name,
		Value: req.Value,
	}
	if err := s.demoRepo.Create(ctx, data); err != nil {
		return nil, err
	}
	return api.Success(mapper.Demo.ToDTO(data), "success"), nil
}
