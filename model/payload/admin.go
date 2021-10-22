package payload

type CreateAdminRequest struct {
	Name     string `json:"name" validate:"required,lte=255"`
	NIM      string `json:"nim" validate:"required,lte=50"`
	Password string `json:"password" validate:"required"`
}

type GetAdminResponse struct {
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	NIM          string `json:"nim"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}
