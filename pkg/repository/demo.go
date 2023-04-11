package repository

import (
	"github.com/vuho-pg/base/pkg/model"
	"gorm.io/gorm"
)

type DemoRepository interface {
	Repository[*model.Demo]
}

type demoRepository struct {
	repository[*model.Demo]
}

func NewDemoRepository(db *gorm.DB) DemoRepository {
	return &demoRepository{repository[*model.Demo]{db}}
}
