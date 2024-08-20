package service

import (
	"context"
	"testing"

	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/repository"
	"github.com/rezaAmiri123/kingscomp/steps/09_game/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountService_CreateOrUpdateWithUserExists(t *testing.T) {
	accRep := &mocks.Account{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{ID: 12, FirstName: "Reza"}, nil,
	).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "Ali"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "Ali", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserNotExists(t *testing.T) {
	accRep := &mocks.Account{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{}, repository.ErrNotFound,
	).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "Ali"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, true, created)
	assert.Equal(t, "Ali", newAcc.FirstName)

	accRep.AssertExpectations(t)
}

func TestAccountService_CreateOrUpdateWithUserHasNotChanged(t *testing.T) {
	accRep := &mocks.Account{}
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 12)).Return(
		entity.Account{ID: 12, FirstName: "Ali"}, nil,
	).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        12,
		FirstName: "Ali",
	})

	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, "Ali", newAcc.FirstName)

	accRep.AssertExpectations(t)
}
