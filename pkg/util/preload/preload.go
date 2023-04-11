package preload

import "gorm.io/gorm"

type PreloadType uint

const (
	PreloadNormal = iota
	PreloadJoin
)

type Applier interface {
	Apply(db *gorm.DB) *gorm.DB
}

type Multiple []Applier

func (p Multiple) Apply(db *gorm.DB) *gorm.DB {
	for _, preload := range p {
		db = preload.Apply(db)
	}
	return db
}

type Single struct {
	pType PreloadType
	value string
	args  []interface{}
}

func (p Single) Apply(db *gorm.DB) *gorm.DB {
	switch p.pType {
	case PreloadJoin:
		return db.Joins(p.value, p.args...)
	case PreloadNormal:
		return db.Preload(p.value, p.args...)
	}
	return db
}

func Many(preloads ...Applier) Applier {
	return Multiple(preloads)
}

func Join(name string, args ...interface{}) Single {
	return Single{
		pType: PreloadJoin,
		value: name,
		args:  args,
	}
}

func Preload(name string, args ...interface{}) Single {
	return Single{
		pType: PreloadNormal,
		value: name,
		args:  args,
	}
}
