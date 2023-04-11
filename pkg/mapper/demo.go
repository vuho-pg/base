package mapper

import (
	"github.com/vuho-pg/base/pkg/dto"
	"github.com/vuho-pg/base/pkg/model"
)

type demoMapper byte

var Demo demoMapper

func (demoMapper) ToDTO(m *model.Demo) *dto.Demo {
	if m == nil {
		return nil
	}
	return &dto.Demo{
		DTO:   Base.ToDTO(m.Model),
		Audit: Audit.ToDTO(m.Audit),
		Name:  m.Name,
		Value: m.Value,
	}
}

func (demoMapper) ToModel(d *dto.Demo) *model.Demo {
	if d == nil {
		return nil
	}
	return &model.Demo{
		Model: Base.ToModel(d.DTO),
		Audit: Audit.ToModel(d.Audit),
		Name:  d.Name,
		Value: d.Value,
	}
}
