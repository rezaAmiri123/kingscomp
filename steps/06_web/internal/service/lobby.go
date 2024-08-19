package service

import "github.com/rezaAmiri123/kingscomp/steps/06_web/internal/repository"

type LobbyService struct {
	Lobby repository.LobbyRepository //todo: refactor, generic common behaviour implementation
}

func NewLobbyService(rep repository.LobbyRepository) *LobbyService {
	return &LobbyService{Lobby: rep}
}
