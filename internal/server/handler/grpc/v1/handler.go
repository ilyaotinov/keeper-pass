package v1

import (
	"context"

	"github.com/ilya372317/pass-keeper/internal/server/dto"
	pb "github.com/ilya372317/pass-keeper/proto"
)

type AuthService interface {
	Login(context.Context, dto.LoginDTO) (string, error)
}

type Handler struct {
	pb.UnimplementedPassServiceServer
	authService AuthService
}

func New(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}
