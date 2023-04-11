package mapper

import (
	"github.com/vuho-pg/base/pkg/dto"
	"github.com/vuho-pg/base/pkg/model"
	"gorm.io/gorm"
	"time"
)

type auditMapper byte

var Audit auditMapper

func (auditMapper) ToDTO(m model.Audit) dto.Audit {
	return dto.Audit{
		CreatedBy: m.CreatedBy,
		UpdatedBy: m.UpdatedBy,
	}
}

func (auditMapper) ToModel(d dto.Audit) model.Audit {
	return model.Audit{
		CreatedBy: d.CreatedBy,
		UpdatedBy: d.UpdatedBy,
	}
}

type baseMapper byte

var Base baseMapper

func (baseMapper) ToDTO(m gorm.Model) dto.DTO {
	return dto.DTO{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Unix(),
		UpdateAt:  m.UpdatedAt.Unix(),
	}
}

func (baseMapper) ToModel(d dto.DTO) gorm.Model {
	return gorm.Model{
		ID:        d.ID,
		CreatedAt: time.Unix(d.CreatedAt, 0),
		UpdatedAt: time.Unix(d.UpdateAt, 0),
	}
}
