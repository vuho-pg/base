package repository

import (
	"context"
	"github.com/vuho-pg/base/pkg/model"
	"github.com/vuho-pg/base/pkg/util/pagination"
	"github.com/vuho-pg/base/pkg/util/preload"
	"gorm.io/gorm"
)

const txKey = "tx"

type Repository[T model.Model] interface {
	DB(ctx context.Context) *gorm.DB
	BeginTx(ctx context.Context, fn func(tx context.Context) error) error
	Create(ctx context.Context, data T) error
	Update(ctx context.Context, data T) error
	Delete(ctx context.Context, data T) error
	FindByID(ctx context.Context, id uint, preloads ...preload.Applier) (T, error)
	FindMany(ctx context.Context, q map[string]interface{}, p pagination.Applier, preloads ...preload.Applier) ([]T, error)
}

type repository[T model.Model] struct {
	db *gorm.DB
}

func (r *repository[T]) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}
	return r.db
}

func (r *repository[T]) SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey, db)
}

func (r *repository[T]) BeginTx(ctx context.Context, fn func(tx context.Context) error) error {
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = r.SetDB(ctx, tx)
		return fn(ctx)
	})
}

func (r *repository[T]) Create(ctx context.Context, data T) error {
	return r.DB(ctx).Model(data).Create(data).Error
}

func (r *repository[T]) Update(ctx context.Context, data T) error {
	return r.DB(ctx).Model(data).Select(data.UpdatableFields()).Updates(data).Error
}

func (r *repository[T]) Delete(ctx context.Context, data T) error {
	return r.DB(ctx).Model(data).Delete(data).Error
}

func (r *repository[T]) FindByID(ctx context.Context, id uint, preloads ...preload.Applier) (T, error) {
	var data T
	db := r.DB(ctx).Model(data)
	db = preload.Many(preloads...).Apply(db)
	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository[T]) FindMany(ctx context.Context, q map[string]interface{}, p pagination.Applier, preloads ...preload.Applier) ([]T, error) {
	res := make([]T, 0)
	var total int64
	db := r.DB(ctx).Model(new(T))
	for k, v := range q {
		db = db.Where(k, v)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}
	p.SetTotal(total)
	db = preload.Many(preloads...).Apply(p.Apply(db))
	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
