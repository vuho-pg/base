package dto

type Demo struct {
	DTO
	Audit
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CreateDemo struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value"`
}
