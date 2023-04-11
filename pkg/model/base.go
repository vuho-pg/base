package model

type AuthorModel struct {
	CreatedBy string
	UpdatedBy string
}

type Model interface {
	UpdatableFields() []string
	TableName() string
}
