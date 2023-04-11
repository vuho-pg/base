package model

type Audit struct {
	CreatedBy string
	UpdatedBy string
}

type Model interface {
	UpdatableFields() []string
	TableName() string
}
