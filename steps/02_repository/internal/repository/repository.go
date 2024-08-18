package repository

import (
	"context"
	"errors"

	"github.com/rezaAmiri123/kingscomp/steps/02_repository/internal/entity"
)

var(
	ErrNotFound = errors.New("entity not found")
)

type CommonBehaviour [T entity.Entity]interface{
	Get(ctx context.Context, id entity.ID)(T, error)
	Save(ctx context.Context, ent entity.Entity)error
}

type AccountRepository interface{
	CommonBehaviour[entity.Account]
}