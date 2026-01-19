package interfaces

import (
	"E-Wallet-wallet/internal/models"
	"context"
)

type IExternal interface {
	ValidateToken(ctx context.Context, token string) (models.TokenData, error)
}
