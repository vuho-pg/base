package dto

type Audit struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

type DTO struct {
	ID        uint  `json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdateAt  int64 `json:"updated_at"`
}
