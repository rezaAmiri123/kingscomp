package service

import "github.com/rezaAmiri123/kingscomp/steps/07_web/internal/repository"

type LobbyService struct {
	Lobby repository.Lobby //todo: refactor, generic common behaviour implementation
}

func NewLobbyService(rep repository.Lobby) *LobbyService {
	return &LobbyService{Lobby: rep}
}
