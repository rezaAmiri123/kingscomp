package gameserver

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rezaAmiri123/kingscomp/steps/14_scale/internal/entity"
	"github.com/rezaAmiri123/kingscomp/steps/14_scale/internal/events"
	"github.com/rezaAmiri123/kingscomp/steps/14_scale/internal/service"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

var (
	ErrGameAlreadyExists = errors.New("game already exists")
	ErrGameNotFound      = errors.New("game not exists")
)

type GameServer struct {
	Config

	games sync.Map
	app   *service.App

	ctx       context.Context
	cancelCtx context.CancelFunc

	pubSub events.PubSub
}

type Config struct {
	ReminderToReadyAfter time.Duration
	ReadyDeadline        time.Duration
	QuestionTimeout      time.Duration
	LobbyAge             time.Duration
	GetReadyDuration     time.Duration
}

func DefaultGameServerConfig() Config {
	return Config{
		ReminderToReadyAfter: DefaultReminderToReadyAfter,
		ReadyDeadline:        DefaultReadyDeadline,
		QuestionTimeout:      DefaultQuestionTimeout,
		LobbyAge:             DefaultLobbyAge,
		GetReadyDuration:     DefaultGetReadyDuration,
	}
}

func NewGameServer(app *service.App, lobbyPubSub events.PubSub, config Config) *GameServer {
	ctx, cancel := context.WithCancel(context.Background())
	gs := &GameServer{
		app:       app,
		Config:    config,
		ctx:       ctx,
		cancelCtx: cancel,
		pubSub:    lobbyPubSub,
	}
	if err := gs.StartupGameServers(context.Background()); err != nil {
		logrus.WithError(err).Errorln("couldn't start up game servers")
	}
	return gs
}

func (g *GameServer) Register(lobbyId string) (*Game, error) {
	game := NewGame(lobbyId, g.app, g, g.Config)
	_, loaded := g.games.LoadOrStore(lobbyId, game)
	if loaded {
		return nil, ErrGameAlreadyExists
	}
	go game.Start(g.ctx)
	return game, nil
}

func (g *GameServer) Game(lobbyId string) (*Game, error) {
	iGame, ok := g.games.Load(lobbyId)
	if !ok {
		return nil, ErrGameNotFound
	}
	return iGame.(*Game), nil
}

func (g *GameServer) MustGame(lobbyId string) *Game {
	game := NewGame(lobbyId, g.app, g, g.Config)
	iGame, ok := g.games.LoadOrStore(lobbyId, game)
	if ok {
		return iGame.(*Game)
	}
	logrus.WithField("lobbyId", lobbyId).Info("Game server was down, restarting the game server")
	go game.Start(g.ctx)
	return game
}

func (g *GameServer) Stop() {
	g.cancelCtx()
}

func (g *GameServer) StartupGameServers(ctx context.Context) error {
	keys, err := g.app.Lobby.AllIDs(ctx, "lobby")
	if err != nil {
		logrus.WithError(err).Errorln("couldn't fetch lobbies")
		return err
	}
	lobbies, err := g.app.Lobby.MGet(ctx, lo.Map[string, entity.ID](keys, func(item string, index int) entity.ID {
		return entity.ID(item)
	})...)

	if err != nil {
		logrus.WithError(err).Errorln("couldn't fetch lobbies to run on startup")
	}

	for _, lobby := range lobbies {
		if lobby.State == "ended" || lobby.CreatedAtUnix < time.Now().Add(-g.LobbyAge).Unix() {
			continue
		}
		g.MustGame(lobby.ID)
		logrus.WithField("lobbyId", lobby.ID).Info("lobby instantiated.")
	}
	return nil
}
