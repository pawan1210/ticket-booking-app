package types

type AddUserRequest struct {
	Name string `json:"name" binding:"required"`
}
