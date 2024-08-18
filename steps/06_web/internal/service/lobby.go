package service

import "github.com/rezaAmiri123/kingscomp/steps/06_web/internal/repository"

type LobbyService struct {
	Lobby repository.LobbyRedisRepository //todo: refactor, generic common behaviour implementation
}

func NewLobbyservice(rep repository.LobbyRedisRepository) *LobbyService {
	return &LobbyService{Lobby: rep}
}
