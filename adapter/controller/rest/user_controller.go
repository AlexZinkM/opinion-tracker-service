package rest

import (
	"encoding/json"
	"net/http"

	"context"
	requests "opinion-tracker-service/adapter/controller/rest/request"
	"opinion-tracker-service/boundry/domain/usecase"
)

type UserController struct {
	CreateUserUseCase usecase.CreateUserUseCase
}

func NewUserController(createUserUseCase usecase.CreateUserUseCase) *UserController {
	return &UserController{CreateUserUseCase: createUserUseCase}
}

// @Summary Create a new user
// @Description Create a new user with the provided email
// @Tags users
// @Accept json
// @Produce json
// @Param user body requests.CreateUserRequest true "User creation request"
// @Success 201 {object} requests.CreateUserReply
// @Router /users [post]
func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(`{"error": "There is only implementation for POST method"}`))
		return
	}

	var req requests.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "invalid request"}`))
		return
	}

	ctx := context.Background()
	dtoReq := req.ToDto()
	dtoReply, err := uc.CreateUserUseCase.CreateUser(ctx, dtoReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "internal error"}`))
		return
	}

	replyJson := requests.FromDto(dtoReply)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(replyJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "failed to encode reply"}`))
		return
	}
}
