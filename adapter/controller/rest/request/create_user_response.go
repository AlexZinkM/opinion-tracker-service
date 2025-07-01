package requests

import (
	"opinion-tracker-service/boundry/dto"
)

type CreateUserRequest struct {
	Email string `json:"email"`
}

func (c *CreateUserRequest) ToDto() dto.CreateUserRequestDTO {
	return dto.CreateUserRequestDTO{
		Email: c.Email,
	}
}
