package service

import (
	"context"
	"errors"
	"time"

	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/06_web/internal/repository"
)

const (
	DefaultState = "home"
)

type AccountService struct {
	Accounts repository.AccountRepository
}

func NewAccountService(rep repository.AccountRepository) *AccountService {
	return &AccountService{Accounts: rep}
}

// CreateOrUpdate creates a new user in the data store or updates the existing user
func (a *AccountService) CreateOrUpdate(ctx context.Context,
	account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.Accounts.Get(ctx, account.EntityID())
	// user exists in the database
	if err == nil {
		if savedAccount.Username != account.Username || savedAccount.FirstName != account.FirstName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			return savedAccount, false, a.Accounts.Save(ctx, savedAccount)
		}
		return savedAccount, false, nil
	}

	// user does not exists in the database
	if errors.Is(err, repository.ErrNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.Accounts.Save(ctx, account)
	}

	return entity.Account{}, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.Accounts.Save(ctx, account)
}
