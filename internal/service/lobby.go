package service

import "github.com/rezaAmiri123/kingscomp/internal/repository"

type LobbyService struct {
	repository.Lobby
}

func NewLobbyService(rep repository.Lobby) *LobbyService {
	return &LobbyService{Lobby: rep}
}
