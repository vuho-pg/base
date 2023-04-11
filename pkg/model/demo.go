package model

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Audit
	Name  string
	Value string
}

func (Demo) TableName() string {
	return "demos"
}

func (Demo) UpdatableFields() []string {
	return []string{
		"name",
		"value",
		"updated_by",
		"created_by",
	}
}
