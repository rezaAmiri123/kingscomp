package service

import "github.com/rezaAmiri123/kingscomp/steps/15_loadbalancer/internal/repository"

type LobbyService struct {
	repository.Lobby //todo: refactor, generic common behaviour implementation
}

func NewLobbyService(rep repository.Lobby) *LobbyService {
	return &LobbyService{Lobby: rep}
}
