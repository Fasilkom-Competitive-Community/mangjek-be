package auth

import (
	"context"

	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Repository interface {
	GetAuthUserFull(ctx context.Context, id string) (uModel.AuthUserFull, error)
}
